// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: cloud/blockstore/config/features.proto

package config

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

type TFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudIds  []string `protobuf:"bytes,1,rep,name=CloudIds" json:"CloudIds,omitempty"`
	FolderIds []string `protobuf:"bytes,2,rep,name=FolderIds" json:"FolderIds,omitempty"`
}

func (x *TFilters) Reset() {
	*x = TFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_config_features_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFilters) ProtoMessage() {}

func (x *TFilters) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_config_features_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFilters.ProtoReflect.Descriptor instead.
func (*TFilters) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_config_features_proto_rawDescGZIP(), []int{0}
}

func (x *TFilters) GetCloudIds() []string {
	if x != nil {
		return x.CloudIds
	}
	return nil
}

func (x *TFilters) GetFolderIds() []string {
	if x != nil {
		return x.FolderIds
	}
	return nil
}

type TFeatureConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Feature name.
	Name *string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// Feature parameters.
	//
	// Types that are assignable to Params:
	//
	//	*TFeatureConfig_Whitelist
	//	*TFeatureConfig_Blacklist
	Params isTFeatureConfig_Params `protobuf_oneof:"Params"`
	// Optional value - for non-binary features.
	Value *string `protobuf:"bytes,4,opt,name=Value" json:"Value,omitempty"`
	// If set, feature can be enabled (for whitelist-features) or disabled
	// (for blacklist-features) even if CloudId and FolderId are not matched by
	// filters - with these probabilities.
	CloudProbability  *float64 `protobuf:"fixed64,5,opt,name=CloudProbability" json:"CloudProbability,omitempty"`
	FolderProbability *float64 `protobuf:"fixed64,6,opt,name=FolderProbability" json:"FolderProbability,omitempty"`
}

func (x *TFeatureConfig) Reset() {
	*x = TFeatureConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_config_features_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFeatureConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFeatureConfig) ProtoMessage() {}

func (x *TFeatureConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_config_features_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFeatureConfig.ProtoReflect.Descriptor instead.
func (*TFeatureConfig) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_config_features_proto_rawDescGZIP(), []int{1}
}

func (x *TFeatureConfig) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (m *TFeatureConfig) GetParams() isTFeatureConfig_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *TFeatureConfig) GetWhitelist() *TFilters {
	if x, ok := x.GetParams().(*TFeatureConfig_Whitelist); ok {
		return x.Whitelist
	}
	return nil
}

func (x *TFeatureConfig) GetBlacklist() *TFilters {
	if x, ok := x.GetParams().(*TFeatureConfig_Blacklist); ok {
		return x.Blacklist
	}
	return nil
}

func (x *TFeatureConfig) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *TFeatureConfig) GetCloudProbability() float64 {
	if x != nil && x.CloudProbability != nil {
		return *x.CloudProbability
	}
	return 0
}

func (x *TFeatureConfig) GetFolderProbability() float64 {
	if x != nil && x.FolderProbability != nil {
		return *x.FolderProbability
	}
	return 0
}

type isTFeatureConfig_Params interface {
	isTFeatureConfig_Params()
}

type TFeatureConfig_Whitelist struct {
	Whitelist *TFilters `protobuf:"bytes,2,opt,name=Whitelist,oneof"`
}

type TFeatureConfig_Blacklist struct {
	Blacklist *TFilters `protobuf:"bytes,3,opt,name=Blacklist,oneof"`
}

func (*TFeatureConfig_Whitelist) isTFeatureConfig_Params() {}

func (*TFeatureConfig_Blacklist) isTFeatureConfig_Params() {}

type TFeaturesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []*TFeatureConfig `protobuf:"bytes,1,rep,name=Features" json:"Features,omitempty"`
}

func (x *TFeaturesConfig) Reset() {
	*x = TFeaturesConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_config_features_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFeaturesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFeaturesConfig) ProtoMessage() {}

func (x *TFeaturesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_config_features_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFeaturesConfig.ProtoReflect.Descriptor instead.
func (*TFeaturesConfig) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_config_features_proto_rawDescGZIP(), []int{2}
}

func (x *TFeaturesConfig) GetFeatures() []*TFeatureConfig {
	if x != nil {
		return x.Features
	}
	return nil
}

var File_cloud_blockstore_config_features_proto protoreflect.FileDescriptor

var file_cloud_blockstore_config_features_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x54, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x0e, 0x54, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x43, 0x0a, 0x09, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x09, 0x57, 0x68, 0x69, 0x74,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x09, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52,
	0x09, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x58, 0x0a, 0x0f, 0x54, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x4e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x2a,
	0x5a, 0x28, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
}

var (
	file_cloud_blockstore_config_features_proto_rawDescOnce sync.Once
	file_cloud_blockstore_config_features_proto_rawDescData = file_cloud_blockstore_config_features_proto_rawDesc
)

func file_cloud_blockstore_config_features_proto_rawDescGZIP() []byte {
	file_cloud_blockstore_config_features_proto_rawDescOnce.Do(func() {
		file_cloud_blockstore_config_features_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_blockstore_config_features_proto_rawDescData)
	})
	return file_cloud_blockstore_config_features_proto_rawDescData
}

var file_cloud_blockstore_config_features_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_blockstore_config_features_proto_goTypes = []interface{}{
	(*TFilters)(nil),        // 0: NCloud.NBlockStore.NProto.TFilters
	(*TFeatureConfig)(nil),  // 1: NCloud.NBlockStore.NProto.TFeatureConfig
	(*TFeaturesConfig)(nil), // 2: NCloud.NBlockStore.NProto.TFeaturesConfig
}
var file_cloud_blockstore_config_features_proto_depIdxs = []int32{
	0, // 0: NCloud.NBlockStore.NProto.TFeatureConfig.Whitelist:type_name -> NCloud.NBlockStore.NProto.TFilters
	0, // 1: NCloud.NBlockStore.NProto.TFeatureConfig.Blacklist:type_name -> NCloud.NBlockStore.NProto.TFilters
	1, // 2: NCloud.NBlockStore.NProto.TFeaturesConfig.Features:type_name -> NCloud.NBlockStore.NProto.TFeatureConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_blockstore_config_features_proto_init() }
func file_cloud_blockstore_config_features_proto_init() {
	if File_cloud_blockstore_config_features_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_blockstore_config_features_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFilters); i {
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
		file_cloud_blockstore_config_features_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFeatureConfig); i {
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
		file_cloud_blockstore_config_features_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFeaturesConfig); i {
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
	file_cloud_blockstore_config_features_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TFeatureConfig_Whitelist)(nil),
		(*TFeatureConfig_Blacklist)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_blockstore_config_features_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_blockstore_config_features_proto_goTypes,
		DependencyIndexes: file_cloud_blockstore_config_features_proto_depIdxs,
		MessageInfos:      file_cloud_blockstore_config_features_proto_msgTypes,
	}.Build()
	File_cloud_blockstore_config_features_proto = out.File
	file_cloud_blockstore_config_features_proto_rawDesc = nil
	file_cloud_blockstore_config_features_proto_goTypes = nil
	file_cloud_blockstore_config_features_proto_depIdxs = nil
}
