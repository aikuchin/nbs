// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/blockstore/private/api/protos/balancer.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EBalancerOpStatus int32

const (
	EBalancerOpStatus_DONT_CHANGE EBalancerOpStatus = 0
	EBalancerOpStatus_ENABLE      EBalancerOpStatus = 1
	EBalancerOpStatus_DISABLE     EBalancerOpStatus = 2
)

// Enum value maps for EBalancerOpStatus.
var (
	EBalancerOpStatus_name = map[int32]string{
		0: "DONT_CHANGE",
		1: "ENABLE",
		2: "DISABLE",
	}
	EBalancerOpStatus_value = map[string]int32{
		"DONT_CHANGE": 0,
		"ENABLE":      1,
		"DISABLE":     2,
	}
)

func (x EBalancerOpStatus) Enum() *EBalancerOpStatus {
	p := new(EBalancerOpStatus)
	*p = x
	return p
}

func (x EBalancerOpStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EBalancerOpStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_blockstore_private_api_protos_balancer_proto_enumTypes[0].Descriptor()
}

func (EBalancerOpStatus) Type() protoreflect.EnumType {
	return &file_cloud_blockstore_private_api_protos_balancer_proto_enumTypes[0]
}

func (x EBalancerOpStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EBalancerOpStatus.Descriptor instead.
func (EBalancerOpStatus) EnumDescriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_balancer_proto_rawDescGZIP(), []int{0}
}

type TConfigureVolumeBalancerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpStatus EBalancerOpStatus `protobuf:"varint,1,opt,name=OpStatus,proto3,enum=NCloud.NBlockStore.NPrivateProto.EBalancerOpStatus" json:"OpStatus,omitempty"`
}

func (x *TConfigureVolumeBalancerRequest) Reset() {
	*x = TConfigureVolumeBalancerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TConfigureVolumeBalancerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TConfigureVolumeBalancerRequest) ProtoMessage() {}

func (x *TConfigureVolumeBalancerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TConfigureVolumeBalancerRequest.ProtoReflect.Descriptor instead.
func (*TConfigureVolumeBalancerRequest) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *TConfigureVolumeBalancerRequest) GetOpStatus() EBalancerOpStatus {
	if x != nil {
		return x.OpStatus
	}
	return EBalancerOpStatus_DONT_CHANGE
}

type TConfigureVolumeBalancerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpStatus EBalancerOpStatus `protobuf:"varint,1,opt,name=OpStatus,proto3,enum=NCloud.NBlockStore.NPrivateProto.EBalancerOpStatus" json:"OpStatus,omitempty"`
}

func (x *TConfigureVolumeBalancerResponse) Reset() {
	*x = TConfigureVolumeBalancerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TConfigureVolumeBalancerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TConfigureVolumeBalancerResponse) ProtoMessage() {}

func (x *TConfigureVolumeBalancerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TConfigureVolumeBalancerResponse.ProtoReflect.Descriptor instead.
func (*TConfigureVolumeBalancerResponse) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_balancer_proto_rawDescGZIP(), []int{1}
}

func (x *TConfigureVolumeBalancerResponse) GetOpStatus() EBalancerOpStatus {
	if x != nil {
		return x.OpStatus
	}
	return EBalancerOpStatus_DONT_CHANGE
}

var File_cloud_blockstore_private_api_protos_balancer_proto protoreflect.FileDescriptor

var file_cloud_blockstore_private_api_protos_balancer_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x1f, 0x54, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x08, 0x4f, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x4e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x4e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x4f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x08, 0x4f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x73, 0x0a, 0x20, 0x54, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x08, 0x4f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x4f, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x4f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a,
	0x3d, 0x0a, 0x11, 0x45, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x4f, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x4f, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x36,
	0x5a, 0x34, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_blockstore_private_api_protos_balancer_proto_rawDescOnce sync.Once
	file_cloud_blockstore_private_api_protos_balancer_proto_rawDescData = file_cloud_blockstore_private_api_protos_balancer_proto_rawDesc
)

func file_cloud_blockstore_private_api_protos_balancer_proto_rawDescGZIP() []byte {
	file_cloud_blockstore_private_api_protos_balancer_proto_rawDescOnce.Do(func() {
		file_cloud_blockstore_private_api_protos_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_blockstore_private_api_protos_balancer_proto_rawDescData)
	})
	return file_cloud_blockstore_private_api_protos_balancer_proto_rawDescData
}

var file_cloud_blockstore_private_api_protos_balancer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_blockstore_private_api_protos_balancer_proto_goTypes = []interface{}{
	(EBalancerOpStatus)(0),                   // 0: NCloud.NBlockStore.NPrivateProto.EBalancerOpStatus
	(*TConfigureVolumeBalancerRequest)(nil),  // 1: NCloud.NBlockStore.NPrivateProto.TConfigureVolumeBalancerRequest
	(*TConfigureVolumeBalancerResponse)(nil), // 2: NCloud.NBlockStore.NPrivateProto.TConfigureVolumeBalancerResponse
}
var file_cloud_blockstore_private_api_protos_balancer_proto_depIdxs = []int32{
	0, // 0: NCloud.NBlockStore.NPrivateProto.TConfigureVolumeBalancerRequest.OpStatus:type_name -> NCloud.NBlockStore.NPrivateProto.EBalancerOpStatus
	0, // 1: NCloud.NBlockStore.NPrivateProto.TConfigureVolumeBalancerResponse.OpStatus:type_name -> NCloud.NBlockStore.NPrivateProto.EBalancerOpStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_blockstore_private_api_protos_balancer_proto_init() }
func file_cloud_blockstore_private_api_protos_balancer_proto_init() {
	if File_cloud_blockstore_private_api_protos_balancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TConfigureVolumeBalancerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TConfigureVolumeBalancerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_blockstore_private_api_protos_balancer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_blockstore_private_api_protos_balancer_proto_goTypes,
		DependencyIndexes: file_cloud_blockstore_private_api_protos_balancer_proto_depIdxs,
		EnumInfos:         file_cloud_blockstore_private_api_protos_balancer_proto_enumTypes,
		MessageInfos:      file_cloud_blockstore_private_api_protos_balancer_proto_msgTypes,
	}.Build()
	File_cloud_blockstore_private_api_protos_balancer_proto = out.File
	file_cloud_blockstore_private_api_protos_balancer_proto_rawDesc = nil
	file_cloud_blockstore_private_api_protos_balancer_proto_goTypes = nil
	file_cloud_blockstore_private_api_protos_balancer_proto_depIdxs = nil
}
