// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: gumball.proto

package __

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gumball_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gumball_proto_msgTypes[0]
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
	return file_gumball_proto_rawDescGZIP(), []int{0}
}

type CountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountResponse) Reset() {
	*x = CountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gumball_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResponse) ProtoMessage() {}

func (x *CountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gumball_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResponse.ProtoReflect.Descriptor instead.
func (*CountResponse) Descriptor() ([]byte, []int) {
	return file_gumball_proto_rawDescGZIP(), []int{1}
}

func (x *CountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *StateResponse) Reset() {
	*x = StateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gumball_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateResponse) ProtoMessage() {}

func (x *StateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gumball_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateResponse.ProtoReflect.Descriptor instead.
func (*StateResponse) Descriptor() ([]byte, []int) {
	return file_gumball_proto_rawDescGZIP(), []int{2}
}

func (x *StateResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type LocatioinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *LocatioinResponse) Reset() {
	*x = LocatioinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gumball_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocatioinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocatioinResponse) ProtoMessage() {}

func (x *LocatioinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gumball_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocatioinResponse.ProtoReflect.Descriptor instead.
func (*LocatioinResponse) Descriptor() ([]byte, []int) {
	return file_gumball_proto_rawDescGZIP(), []int{3}
}

func (x *LocatioinResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_gumball_proto protoreflect.FileDescriptor

var file_gumball_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x75, 0x6d, 0x62, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x25, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2f, 0x0a,
	0x11, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa0,
	0x01, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gumball_proto_rawDescOnce sync.Once
	file_gumball_proto_rawDescData = file_gumball_proto_rawDesc
)

func file_gumball_proto_rawDescGZIP() []byte {
	file_gumball_proto_rawDescOnce.Do(func() {
		file_gumball_proto_rawDescData = protoimpl.X.CompressGZIP(file_gumball_proto_rawDescData)
	})
	return file_gumball_proto_rawDescData
}

var file_gumball_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gumball_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: proto.Empty
	(*CountResponse)(nil),     // 1: proto.CountResponse
	(*StateResponse)(nil),     // 2: proto.StateResponse
	(*LocatioinResponse)(nil), // 3: proto.LocatioinResponse
}
var file_gumball_proto_depIdxs = []int32{
	0, // 0: proto.Monitor.GetCount:input_type -> proto.Empty
	0, // 1: proto.Monitor.GetState:input_type -> proto.Empty
	0, // 2: proto.Monitor.GetLocation:input_type -> proto.Empty
	1, // 3: proto.Monitor.GetCount:output_type -> proto.CountResponse
	2, // 4: proto.Monitor.GetState:output_type -> proto.StateResponse
	3, // 5: proto.Monitor.GetLocation:output_type -> proto.LocatioinResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gumball_proto_init() }
func file_gumball_proto_init() {
	if File_gumball_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gumball_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gumball_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountResponse); i {
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
		file_gumball_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateResponse); i {
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
		file_gumball_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocatioinResponse); i {
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
			RawDescriptor: file_gumball_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gumball_proto_goTypes,
		DependencyIndexes: file_gumball_proto_depIdxs,
		MessageInfos:      file_gumball_proto_msgTypes,
	}.Build()
	File_gumball_proto = out.File
	file_gumball_proto_rawDesc = nil
	file_gumball_proto_goTypes = nil
	file_gumball_proto_depIdxs = nil
}
