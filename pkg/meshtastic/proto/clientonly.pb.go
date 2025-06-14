// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: meshtastic/clientonly.proto

package proto

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

// This abstraction is used to contain any configuration for provisioning a node on any client.
// It is useful for importing and exporting configurations.
type DeviceProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Long name for the node
	LongName *string `protobuf:"bytes,1,opt,name=long_name,json=longName,proto3,oneof" json:"long_name,omitempty"`
	// Short name of the node
	ShortName *string `protobuf:"bytes,2,opt,name=short_name,json=shortName,proto3,oneof" json:"short_name,omitempty"`
	// The url of the channels from our node
	ChannelUrl *string `protobuf:"bytes,3,opt,name=channel_url,json=channelUrl,proto3,oneof" json:"channel_url,omitempty"`
	// The Config of the node
	Config *LocalConfig `protobuf:"bytes,4,opt,name=config,proto3,oneof" json:"config,omitempty"`
	// The ModuleConfig of the node
	ModuleConfig *LocalModuleConfig `protobuf:"bytes,5,opt,name=module_config,json=moduleConfig,proto3,oneof" json:"module_config,omitempty"`
	// Fixed position data
	FixedPosition *Position `protobuf:"bytes,6,opt,name=fixed_position,json=fixedPosition,proto3,oneof" json:"fixed_position,omitempty"`
	// Ringtone for ExternalNotification
	Ringtone *string `protobuf:"bytes,7,opt,name=ringtone,proto3,oneof" json:"ringtone,omitempty"`
	// Predefined messages for CannedMessage
	CannedMessages *string `protobuf:"bytes,8,opt,name=canned_messages,json=cannedMessages,proto3,oneof" json:"canned_messages,omitempty"`
}

func (x *DeviceProfile) Reset() {
	*x = DeviceProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_clientonly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceProfile) ProtoMessage() {}

func (x *DeviceProfile) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_clientonly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceProfile.ProtoReflect.Descriptor instead.
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return file_meshtastic_clientonly_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceProfile) GetLongName() string {
	if x != nil && x.LongName != nil {
		return *x.LongName
	}
	return ""
}

func (x *DeviceProfile) GetShortName() string {
	if x != nil && x.ShortName != nil {
		return *x.ShortName
	}
	return ""
}

func (x *DeviceProfile) GetChannelUrl() string {
	if x != nil && x.ChannelUrl != nil {
		return *x.ChannelUrl
	}
	return ""
}

func (x *DeviceProfile) GetConfig() *LocalConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *DeviceProfile) GetModuleConfig() *LocalModuleConfig {
	if x != nil {
		return x.ModuleConfig
	}
	return nil
}

func (x *DeviceProfile) GetFixedPosition() *Position {
	if x != nil {
		return x.FixedPosition
	}
	return nil
}

func (x *DeviceProfile) GetRingtone() string {
	if x != nil && x.Ringtone != nil {
		return *x.Ringtone
	}
	return ""
}

func (x *DeviceProfile) GetCannedMessages() string {
	if x != nil && x.CannedMessages != nil {
		return *x.CannedMessages
	}
	return ""
}

var File_meshtastic_clientonly_proto protoreflect.FileDescriptor

var file_meshtastic_clientonly_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x68, 0x74,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6f, 0x6e, 0x6c, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x04, 0x0a,
	0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01,
	0x12, 0x47, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x04, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0e, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x05, 0x52, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72,
	0x69, 0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x08, 0x72, 0x69, 0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c,
	0x6f, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x69, 0x6e, 0x67,
	0x74, 0x6f, 0x6e, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x53, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42,
	0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x5a, 0x10, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x14, 0x4d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_clientonly_proto_rawDescOnce sync.Once
	file_meshtastic_clientonly_proto_rawDescData = file_meshtastic_clientonly_proto_rawDesc
)

func file_meshtastic_clientonly_proto_rawDescGZIP() []byte {
	file_meshtastic_clientonly_proto_rawDescOnce.Do(func() {
		file_meshtastic_clientonly_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_clientonly_proto_rawDescData)
	})
	return file_meshtastic_clientonly_proto_rawDescData
}

var file_meshtastic_clientonly_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meshtastic_clientonly_proto_goTypes = []interface{}{
	(*DeviceProfile)(nil),     // 0: meshtastic.DeviceProfile
	(*LocalConfig)(nil),       // 1: meshtastic.LocalConfig
	(*LocalModuleConfig)(nil), // 2: meshtastic.LocalModuleConfig
	(*Position)(nil),          // 3: meshtastic.Position
}
var file_meshtastic_clientonly_proto_depIdxs = []int32{
	1, // 0: meshtastic.DeviceProfile.config:type_name -> meshtastic.LocalConfig
	2, // 1: meshtastic.DeviceProfile.module_config:type_name -> meshtastic.LocalModuleConfig
	3, // 2: meshtastic.DeviceProfile.fixed_position:type_name -> meshtastic.Position
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_meshtastic_clientonly_proto_init() }
func file_meshtastic_clientonly_proto_init() {
	if File_meshtastic_clientonly_proto != nil {
		return
	}
	file_meshtastic_localonly_proto_init()
	file_meshtastic_mesh_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meshtastic_clientonly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceProfile); i {
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
	file_meshtastic_clientonly_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meshtastic_clientonly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_clientonly_proto_goTypes,
		DependencyIndexes: file_meshtastic_clientonly_proto_depIdxs,
		MessageInfos:      file_meshtastic_clientonly_proto_msgTypes,
	}.Build()
	File_meshtastic_clientonly_proto = out.File
	file_meshtastic_clientonly_proto_rawDesc = nil
	file_meshtastic_clientonly_proto_goTypes = nil
	file_meshtastic_clientonly_proto_depIdxs = nil
}
