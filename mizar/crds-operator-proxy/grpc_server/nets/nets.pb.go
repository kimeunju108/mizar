// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: nets/nets.proto

package nets

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// requests
type Net struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip             string `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Prefix         string `protobuf:"bytes,2,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	Vni            string `protobuf:"bytes,3,opt,name=Vni,proto3" json:"Vni,omitempty"`
	Vpc            string `protobuf:"bytes,4,opt,name=Vpc,proto3" json:"Vpc,omitempty"`
	Status         string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Bouncers       string `protobuf:"bytes,6,opt,name=Bouncers,proto3" json:"Bouncers,omitempty"`
	CreateTime     string `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	ProvisionDelay string `protobuf:"bytes,8,opt,name=ProvisionDelay,proto3" json:"ProvisionDelay,omitempty"`
}

func (x *Net) Reset() {
	*x = Net{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nets_nets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Net) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Net) ProtoMessage() {}

func (x *Net) ProtoReflect() protoreflect.Message {
	mi := &file_nets_nets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Net.ProtoReflect.Descriptor instead.
func (*Net) Descriptor() ([]byte, []int) {
	return file_nets_nets_proto_rawDescGZIP(), []int{0}
}

func (x *Net) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Net) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Net) GetVni() string {
	if x != nil {
		return x.Vni
	}
	return ""
}

func (x *Net) GetVpc() string {
	if x != nil {
		return x.Vpc
	}
	return ""
}

func (x *Net) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Net) GetBouncers() string {
	if x != nil {
		return x.Bouncers
	}
	return ""
}

func (x *Net) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Net) GetProvisionDelay() string {
	if x != nil {
		return x.ProvisionDelay
	}
	return ""
}

type NetId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *NetId) Reset() {
	*x = NetId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nets_nets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetId) ProtoMessage() {}

func (x *NetId) ProtoReflect() protoreflect.Message {
	mi := &file_nets_nets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetId.ProtoReflect.Descriptor instead.
func (*NetId) Descriptor() ([]byte, []int) {
	return file_nets_nets_proto_rawDescGZIP(), []int{1}
}

func (x *NetId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

//response
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nets_nets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_nets_nets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_nets_nets_proto_rawDescGZIP(), []int{2}
}

type NetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonReturncode string `protobuf:"bytes,1,opt,name=json_returncode,json=jsonReturncode,proto3" json:"json_returncode,omitempty"`
}

func (x *NetsResponse) Reset() {
	*x = NetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nets_nets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetsResponse) ProtoMessage() {}

func (x *NetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nets_nets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetsResponse.ProtoReflect.Descriptor instead.
func (*NetsResponse) Descriptor() ([]byte, []int) {
	return file_nets_nets_proto_rawDescGZIP(), []int{3}
}

func (x *NetsResponse) GetJsonReturncode() string {
	if x != nil {
		return x.JsonReturncode
	}
	return ""
}

var File_nets_nets_proto protoreflect.FileDescriptor

var file_nets_nets_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6e, 0x65, 0x74, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x03, 0x4e, 0x65, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x6e, 0x69, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x56, 0x6e, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x70, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x56, 0x70, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x72, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x17, 0x0a, 0x05, 0x4e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x37, 0x0a, 0x0c, 0x4e, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0xdb, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x12,
	0x09, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x1a, 0x0b, 0x2e, 0x6e, 0x65, 0x74,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x12, 0x09, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x4e, 0x65,
	0x74, 0x1a, 0x0b, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x2c, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x6e, 0x65,
	0x74, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e,
	0x4e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x6e, 0x65,
	0x74, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x4e, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x49,
	0x64, 0x1a, 0x0b, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x5c, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x63, 0x72, 0x64, 0x73, 0x2d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x6e, 0x65, 0x74, 0x73, 0x42, 0x09, 0x4e, 0x65, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x29, 0x6d, 0x69, 0x7a, 0x61, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x64, 0x73, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nets_nets_proto_rawDescOnce sync.Once
	file_nets_nets_proto_rawDescData = file_nets_nets_proto_rawDesc
)

func file_nets_nets_proto_rawDescGZIP() []byte {
	file_nets_nets_proto_rawDescOnce.Do(func() {
		file_nets_nets_proto_rawDescData = protoimpl.X.CompressGZIP(file_nets_nets_proto_rawDescData)
	})
	return file_nets_nets_proto_rawDescData
}

var file_nets_nets_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nets_nets_proto_goTypes = []interface{}{
	(*Net)(nil),          // 0: nets.Net
	(*NetId)(nil),        // 1: nets.NetId
	(*Empty)(nil),        // 2: nets.Empty
	(*NetsResponse)(nil), // 3: nets.NetsResponse
}
var file_nets_nets_proto_depIdxs = []int32{
	0, // 0: nets.NetsService.CreateNet:input_type -> nets.Net
	0, // 1: nets.NetsService.UpdateNet:input_type -> nets.Net
	1, // 2: nets.NetsService.ReadNet:input_type -> nets.NetId
	1, // 3: nets.NetsService.DeleteNet:input_type -> nets.NetId
	1, // 4: nets.NetsService.ResumeNet:input_type -> nets.NetId
	2, // 5: nets.NetsService.CreateNet:output_type -> nets.Empty
	2, // 6: nets.NetsService.UpdateNet:output_type -> nets.Empty
	3, // 7: nets.NetsService.ReadNet:output_type -> nets.NetsResponse
	2, // 8: nets.NetsService.DeleteNet:output_type -> nets.Empty
	2, // 9: nets.NetsService.ResumeNet:output_type -> nets.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nets_nets_proto_init() }
func file_nets_nets_proto_init() {
	if File_nets_nets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nets_nets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Net); i {
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
		file_nets_nets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetId); i {
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
		file_nets_nets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_nets_nets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetsResponse); i {
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
			RawDescriptor: file_nets_nets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nets_nets_proto_goTypes,
		DependencyIndexes: file_nets_nets_proto_depIdxs,
		MessageInfos:      file_nets_nets_proto_msgTypes,
	}.Build()
	File_nets_nets_proto = out.File
	file_nets_nets_proto_rawDesc = nil
	file_nets_nets_proto_goTypes = nil
	file_nets_nets_proto_depIdxs = nil
}
