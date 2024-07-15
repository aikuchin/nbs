package disks

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	disk_manager "github.com/ydb-platform/nbs/cloud/disk_manager/api"
	"github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/clients/nbs"
	dataplane_protos "github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/dataplane/protos"
	"github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/performance"
	performance_config "github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/performance/config"
	"github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/resources"
	"github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/services/disks/protos"
	"github.com/ydb-platform/nbs/cloud/disk_manager/internal/pkg/types"
	"github.com/ydb-platform/nbs/cloud/tasks"
	"github.com/ydb-platform/nbs/cloud/tasks/errors"
	"github.com/ydb-platform/nbs/cloud/tasks/headers"
)

////////////////////////////////////////////////////////////////////////////////

type createDiskFromImageTask struct {
	performanceConfig *performance_config.PerformanceConfig
	storage           resources.Storage
	scheduler         tasks.Scheduler
	nbsFactory        nbs.Factory
	request           *protos.CreateDiskFromImageRequest
	state             *protos.CreateDiskFromImageTaskState
}

func (t *createDiskFromImageTask) Save() ([]byte, error) {
	return proto.Marshal(t.state)
}

func (t *createDiskFromImageTask) Load(request, state []byte) error {
	t.request = &protos.CreateDiskFromImageRequest{}
	err := proto.Unmarshal(request, t.request)
	if err != nil {
		return err
	}

	t.state = &protos.CreateDiskFromImageTaskState{}
	return proto.Unmarshal(state, t.state)
}

func (t *createDiskFromImageTask) Run(
	ctx context.Context,
	execCtx tasks.ExecutionContext,
) error {

	params := t.request.Params

	client, err := t.nbsFactory.GetClient(ctx, params.Disk.ZoneId)
	if err != nil {
		return err
	}

	selfTaskID := execCtx.GetTaskID()

	diskMeta, err := t.storage.CreateDisk(ctx, resources.DiskMeta{
		ID:          params.Disk.DiskId,
		ZoneID:      params.Disk.ZoneId,
		SrcImageID:  t.request.SrcImageId,
		BlocksCount: params.BlocksCount,
		BlockSize:   params.BlockSize,
		Kind:        diskKindToString(params.Kind),
		CloudID:     params.CloudId,
		FolderID:    params.FolderId,

		CreateRequest: t.request,
		CreateTaskID:  selfTaskID,
		CreatingAt:    time.Now(),
		CreatedBy:     "", // TODO: Extract CreatedBy from execCtx
	})
	if err != nil {
		return err
	}

	if diskMeta == nil {
		return errors.NewNonCancellableErrorf(
			"id %v is not accepted",
			params.Disk.DiskId,
		)
	}

	imageMeta, err := t.storage.GetImageMeta(ctx, t.request.SrcImageId)
	if err != nil {
		return err
	}

	diskEncryption := types.EncryptionMode_NO_ENCRYPTION
	imageEncryption := types.EncryptionMode_NO_ENCRYPTION

	if params.EncryptionDesc != nil {
		diskEncryption = params.EncryptionDesc.Mode
	}

	if imageMeta != nil {
		execCtx.SetEstimate(performance.Estimate(
			imageMeta.StorageSize,
			t.performanceConfig.GetCreateDiskFromImageBandwidthMiBs(),
		))

		if imageMeta.Encryption != nil {
			imageEncryption = imageMeta.Encryption.Mode
		}
	}

	if imageEncryption != types.EncryptionMode_NO_ENCRYPTION &&
		diskEncryption != imageEncryption {

		message := fmt.Sprintf(
			"encryption mode should be the same for disk (%v) and encrypted image (%v)",
			diskEncryption,
			imageEncryption,
		)
		if imageMeta == nil {
			return errors.NewSilentNonRetriableErrorf(message)
		}
		return errors.NewNonRetriableErrorf(message)
	}

	encryption := params.EncryptionDesc
	if imageEncryption != types.EncryptionMode_NO_ENCRYPTION {
		encryption = imageMeta.Encryption
	}

	err = client.Create(ctx, nbs.CreateDiskParams{
		ID:                      params.Disk.DiskId,
		BlocksCount:             params.BlocksCount,
		BlockSize:               params.BlockSize,
		Kind:                    params.Kind,
		CloudID:                 params.CloudId,
		FolderID:                params.FolderId,
		TabletVersion:           params.TabletVersion,
		PlacementGroupID:        params.PlacementGroupId,
		PlacementPartitionIndex: params.PlacementPartitionIndex,
		StoragePoolName:         params.StoragePoolName,
		AgentIds:                params.AgentIds,
		EncryptionDesc:          encryption,
	})
	if err != nil {
		return err
	}

	var taskID string

	// Old images without metadata we consider as not dataplane.
	if imageMeta != nil && imageMeta.UseDataplaneTasks {
		taskID, err = t.scheduler.ScheduleZonalTask(
			headers.SetIncomingIdempotencyKey(ctx, selfTaskID),
			"dataplane.TransferFromSnapshotToDisk",
			"",
			params.Disk.ZoneId,
			&dataplane_protos.TransferFromSnapshotToDiskRequest{
				SrcSnapshotId: t.request.SrcImageId,
				DstDisk:       params.Disk,
				DstEncryption: encryption,
			},
		)

		t.state.DataplaneTaskId = taskID
	} else {
		taskID, err = t.scheduler.ScheduleZonalTask(
			headers.SetIncomingIdempotencyKey(ctx, selfTaskID),
			"dataplane.TransferFromLegacySnapshotToDisk",
			"",
			params.Disk.ZoneId,
			&dataplane_protos.TransferFromSnapshotToDiskRequest{
				SrcSnapshotId: t.request.SrcImageId,
				DstDisk:       params.Disk,
				DstEncryption: encryption,
			},
		)

		t.state.DataplaneTaskId = taskID
	}
	if err != nil {
		return err
	}

	err = execCtx.SaveState(ctx)
	if err != nil {
		return err
	}

	_, err = t.scheduler.WaitTask(ctx, execCtx, taskID)
	if err != nil {
		return err
	}

	diskMeta.CreatedAt = time.Now()
	return t.storage.DiskCreated(ctx, *diskMeta)
}

func (t *createDiskFromImageTask) Cancel(
	ctx context.Context,
	execCtx tasks.ExecutionContext,
) error {

	params := t.request.Params

	client, err := t.nbsFactory.GetClient(ctx, params.Disk.ZoneId)
	if err != nil {
		return err
	}

	selfTaskID := execCtx.GetTaskID()

	disk, err := t.storage.DeleteDisk(
		ctx,
		params.Disk.DiskId,
		selfTaskID,
		time.Now(),
	)
	if err != nil {
		return err
	}

	if disk == nil {
		return errors.NewNonCancellableErrorf(
			"id %v is not accepted",
			params.Disk.DiskId,
		)
	}

	err = client.Delete(ctx, params.Disk.DiskId)
	if err != nil {
		return err
	}

	return t.storage.DiskDeleted(ctx, params.Disk.DiskId, time.Now())
}

func (t *createDiskFromImageTask) GetMetadata(
	ctx context.Context,
) (proto.Message, error) {

	metadata := &disk_manager.CreateDiskMetadata{}

	if len(t.state.DataplaneTaskId) != 0 {
		message, err := t.scheduler.GetTaskMetadata(
			ctx,
			t.state.DataplaneTaskId,
		)
		if err != nil {
			return nil, err
		}

		transferMetadata, ok := message.(*dataplane_protos.TransferFromSnapshotToDiskMetadata)
		if ok {
			metadata.Progress = transferMetadata.Progress
		}
	}

	return metadata, nil
}

func (t *createDiskFromImageTask) GetResponse() proto.Message {
	return &empty.Empty{}
}
