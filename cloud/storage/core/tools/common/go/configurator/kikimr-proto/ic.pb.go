// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/storage/core/tools/common/go/configurator/kikimr-proto/ic.proto

package kikimr_proto

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

type TDuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds      *uint64 `protobuf:"varint,1,opt,name=Seconds" json:"Seconds,omitempty"`
	Milliseconds *uint64 `protobuf:"varint,2,opt,name=Milliseconds" json:"Milliseconds,omitempty"`
	Microseconds *uint64 `protobuf:"varint,3,opt,name=Microseconds" json:"Microseconds,omitempty"`
}

func (x *TDuration) Reset() {
	*x = TDuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDuration) ProtoMessage() {}

func (x *TDuration) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDuration.ProtoReflect.Descriptor instead.
func (*TDuration) Descriptor() ([]byte, []int) {
	return file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescGZIP(), []int{0}
}

func (x *TDuration) GetSeconds() uint64 {
	if x != nil && x.Seconds != nil {
		return *x.Seconds
	}
	return 0
}

func (x *TDuration) GetMilliseconds() uint64 {
	if x != nil && x.Milliseconds != nil {
		return *x.Milliseconds
	}
	return 0
}

func (x *TDuration) GetMicroseconds() uint64 {
	if x != nil && x.Microseconds != nil {
		return *x.Microseconds
	}
	return 0
}

type TInterconnectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTcp                    *bool      `protobuf:"varint,1,opt,name=StartTcp" json:"StartTcp,omitempty"`
	MaxInflightAmountOfDataInKB *uint32    `protobuf:"varint,2,opt,name=MaxInflightAmountOfDataInKB" json:"MaxInflightAmountOfDataInKB,omitempty"`
	MergePerPeerCounters        *bool      `protobuf:"varint,3,opt,name=MergePerPeerCounters" json:"MergePerPeerCounters,omitempty"`
	HandshakeTimeoutDuration    *TDuration `protobuf:"bytes,4,opt,name=HandshakeTimeoutDuration" json:"HandshakeTimeoutDuration,omitempty"`
}

func (x *TInterconnectConfig) Reset() {
	*x = TInterconnectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TInterconnectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TInterconnectConfig) ProtoMessage() {}

func (x *TInterconnectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TInterconnectConfig.ProtoReflect.Descriptor instead.
func (*TInterconnectConfig) Descriptor() ([]byte, []int) {
	return file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescGZIP(), []int{1}
}

func (x *TInterconnectConfig) GetStartTcp() bool {
	if x != nil && x.StartTcp != nil {
		return *x.StartTcp
	}
	return false
}

func (x *TInterconnectConfig) GetMaxInflightAmountOfDataInKB() uint32 {
	if x != nil && x.MaxInflightAmountOfDataInKB != nil {
		return *x.MaxInflightAmountOfDataInKB
	}
	return 0
}

func (x *TInterconnectConfig) GetMergePerPeerCounters() bool {
	if x != nil && x.MergePerPeerCounters != nil {
		return *x.MergePerPeerCounters
	}
	return false
}

func (x *TInterconnectConfig) GetHandshakeTimeoutDuration() *TDuration {
	if x != nil {
		return x.HandshakeTimeoutDuration
	}
	return nil
}

var File_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto protoreflect.FileDescriptor

var file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x6b, 0x69, 0x6b, 0x69, 0x6d, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x09, 0x54, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xef, 0x01, 0x0a, 0x13, 0x54, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x63, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x63, 0x70, 0x12, 0x40, 0x0a, 0x1b, 0x4d, 0x61,
	0x78, 0x49, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f,
	0x66, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x4b, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x1b, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x4f, 0x66, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x4b, 0x42, 0x12, 0x32, 0x0a, 0x14,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x50, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x50, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x46, 0x0a, 0x18, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4f, 0x5a, 0x4d, 0x61, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6b, 0x69, 0x6b,
	0x69, 0x6d, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescOnce sync.Once
	file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescData = file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDesc
)

func file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescGZIP() []byte {
	file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescOnce.Do(func() {
		file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescData)
	})
	return file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDescData
}

var file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_goTypes = []interface{}{
	(*TDuration)(nil),           // 0: TDuration
	(*TInterconnectConfig)(nil), // 1: TInterconnectConfig
}
var file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_depIdxs = []int32{
	0, // 0: TInterconnectConfig.HandshakeTimeoutDuration:type_name -> TDuration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_init() }
func file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_init() {
	if File_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDuration); i {
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
		file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TInterconnectConfig); i {
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
			RawDescriptor: file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_goTypes,
		DependencyIndexes: file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_depIdxs,
		MessageInfos:      file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_msgTypes,
	}.Build()
	File_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto = out.File
	file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_rawDesc = nil
	file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_goTypes = nil
	file_cloud_storage_core_tools_common_go_configurator_kikimr_proto_ic_proto_depIdxs = nil
}
