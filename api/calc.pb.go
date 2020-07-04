// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: api/calc.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AdditionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *AdditionMessage) Reset() {
	*x = AdditionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionMessage) ProtoMessage() {}

func (x *AdditionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionMessage.ProtoReflect.Descriptor instead.
func (*AdditionMessage) Descriptor() ([]byte, []int) {
	return file_api_calc_proto_rawDescGZIP(), []int{0}
}

func (x *AdditionMessage) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *AdditionMessage) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_api_calc_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AdditionResultMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResNum int32  `protobuf:"varint,1,opt,name=resNum,proto3" json:"resNum,omitempty"`
	Error  *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AdditionResultMessage) Reset() {
	*x = AdditionResultMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionResultMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionResultMessage) ProtoMessage() {}

func (x *AdditionResultMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionResultMessage.ProtoReflect.Descriptor instead.
func (*AdditionResultMessage) Descriptor() ([]byte, []int) {
	return file_api_calc_proto_rawDescGZIP(), []int{2}
}

func (x *AdditionResultMessage) GetResNum() int32 {
	if x != nil {
		return x.ResNum
	}
	return 0
}

func (x *AdditionResultMessage) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_api_calc_proto protoreflect.FileDescriptor

var file_api_calc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x39, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32,
	0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x4b, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_calc_proto_rawDescOnce sync.Once
	file_api_calc_proto_rawDescData = file_api_calc_proto_rawDesc
)

func file_api_calc_proto_rawDescGZIP() []byte {
	file_api_calc_proto_rawDescOnce.Do(func() {
		file_api_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_calc_proto_rawDescData)
	})
	return file_api_calc_proto_rawDescData
}

var file_api_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_calc_proto_goTypes = []interface{}{
	(*AdditionMessage)(nil),       // 0: api.AdditionMessage
	(*Error)(nil),                 // 1: api.Error
	(*AdditionResultMessage)(nil), // 2: api.AdditionResultMessage
}
var file_api_calc_proto_depIdxs = []int32{
	1, // 0: api.AdditionResultMessage.error:type_name -> api.Error
	0, // 1: api.Addition.AddNumber:input_type -> api.AdditionMessage
	2, // 2: api.Addition.AddNumber:output_type -> api.AdditionResultMessage
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_calc_proto_init() }
func file_api_calc_proto_init() {
	if File_api_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionMessage); i {
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
		file_api_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_api_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionResultMessage); i {
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
			RawDescriptor: file_api_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calc_proto_goTypes,
		DependencyIndexes: file_api_calc_proto_depIdxs,
		MessageInfos:      file_api_calc_proto_msgTypes,
	}.Build()
	File_api_calc_proto = out.File
	file_api_calc_proto_rawDesc = nil
	file_api_calc_proto_goTypes = nil
	file_api_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdditionClient is the client API for Addition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionClient interface {
	AddNumber(ctx context.Context, in *AdditionMessage, opts ...grpc.CallOption) (*AdditionResultMessage, error)
}

type additionClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionClient(cc grpc.ClientConnInterface) AdditionClient {
	return &additionClient{cc}
}

func (c *additionClient) AddNumber(ctx context.Context, in *AdditionMessage, opts ...grpc.CallOption) (*AdditionResultMessage, error) {
	out := new(AdditionResultMessage)
	err := c.cc.Invoke(ctx, "/api.Addition/AddNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionServer is the server API for Addition service.
type AdditionServer interface {
	AddNumber(context.Context, *AdditionMessage) (*AdditionResultMessage, error)
}

// UnimplementedAdditionServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionServer struct {
}

func (*UnimplementedAdditionServer) AddNumber(context.Context, *AdditionMessage) (*AdditionResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNumber not implemented")
}

func RegisterAdditionServer(s *grpc.Server, srv AdditionServer) {
	s.RegisterService(&_Addition_serviceDesc, srv)
}

func _Addition_AddNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdditionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServer).AddNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Addition/AddNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServer).AddNumber(ctx, req.(*AdditionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Addition",
	HandlerType: (*AdditionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumber",
			Handler:    _Addition_AddNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/calc.proto",
}
