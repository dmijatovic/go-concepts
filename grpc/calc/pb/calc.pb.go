// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: pb/calc.proto

package calc

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

type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num_1     int32  `protobuf:"varint,1,opt,name=num_1,json=num1,proto3" json:"num_1,omitempty"`
	Num_2     int32  `protobuf:"varint,2,opt,name=num_2,json=num2,proto3" json:"num_2,omitempty"`
	Operation string `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_pb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetNum_1() int32 {
	if x != nil {
		return x.Num_1
	}
	return 0
}

func (x *CalcRequest) GetNum_2() int32 {
	if x != nil {
		return x.Num_2
	}
	return 0
}

func (x *CalcRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

type CalcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalcResponse) Reset() {
	*x = CalcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcResponse) ProtoMessage() {}

func (x *CalcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcResponse.ProtoReflect.Descriptor instead.
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return file_pb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_pb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *AverageRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mean float64 `protobuf:"fixed64,1,opt,name=mean,proto3" json:"mean,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_pb_calc_proto_rawDescGZIP(), []int{3}
}

func (x *AverageResponse) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

var File_pb_calc_proto protoreflect.FileDescriptor

var file_pb_calc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x55, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x5f, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x75, 0x6d,
	0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0c,
	0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x22, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x32,
	0x94, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x69, 0x76,
	0x69, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x62, 0x3b, 0x63, 0x61, 0x6c,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_calc_proto_rawDescOnce sync.Once
	file_pb_calc_proto_rawDescData = file_pb_calc_proto_rawDesc
)

func file_pb_calc_proto_rawDescGZIP() []byte {
	file_pb_calc_proto_rawDescOnce.Do(func() {
		file_pb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_calc_proto_rawDescData)
	})
	return file_pb_calc_proto_rawDescData
}

var file_pb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_calc_proto_goTypes = []interface{}{
	(*CalcRequest)(nil),     // 0: calc.CalcRequest
	(*CalcResponse)(nil),    // 1: calc.CalcResponse
	(*AverageRequest)(nil),  // 2: calc.AverageRequest
	(*AverageResponse)(nil), // 3: calc.AverageResponse
}
var file_pb_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalcService.Add:input_type -> calc.CalcRequest
	0, // 1: calc.CalcService.Subtract:input_type -> calc.CalcRequest
	0, // 2: calc.CalcService.Multiply:input_type -> calc.CalcRequest
	0, // 3: calc.CalcService.Divide:input_type -> calc.CalcRequest
	2, // 4: calc.CalcService.Average:input_type -> calc.AverageRequest
	1, // 5: calc.CalcService.Add:output_type -> calc.CalcResponse
	1, // 6: calc.CalcService.Subtract:output_type -> calc.CalcResponse
	1, // 7: calc.CalcService.Multiply:output_type -> calc.CalcResponse
	1, // 8: calc.CalcService.Divide:output_type -> calc.CalcResponse
	3, // 9: calc.CalcService.Average:output_type -> calc.AverageResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_calc_proto_init() }
func file_pb_calc_proto_init() {
	if File_pb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
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
		file_pb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcResponse); i {
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
		file_pb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_pb_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
			RawDescriptor: file_pb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_calc_proto_goTypes,
		DependencyIndexes: file_pb_calc_proto_depIdxs,
		MessageInfos:      file_pb_calc_proto_msgTypes,
	}.Build()
	File_pb_calc_proto = out.File
	file_pb_calc_proto_rawDesc = nil
	file_pb_calc_proto_goTypes = nil
	file_pb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	Add(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Subtract(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Multiply(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Divide(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	// client side stram
	Average(ctx context.Context, opts ...grpc.CallOption) (CalcService_AverageClient, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Add(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Subtract(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Multiply(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Divide(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalcService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calc.CalcService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceAverageClient{stream}
	return x, nil
}

type CalcService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calcServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	Add(context.Context, *CalcRequest) (*CalcResponse, error)
	Subtract(context.Context, *CalcRequest) (*CalcResponse, error)
	Multiply(context.Context, *CalcRequest) (*CalcResponse, error)
	Divide(context.Context, *CalcRequest) (*CalcResponse, error)
	// client side stram
	Average(CalcService_AverageServer) error
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Add(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalcServiceServer) Subtract(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (*UnimplementedCalcServiceServer) Multiply(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedCalcServiceServer) Divide(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (*UnimplementedCalcServiceServer) Average(CalcService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Add(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Subtract(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Multiply(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Divide(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).Average(&calcServiceAverageServer{stream})
}

type CalcService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calcServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalcService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _CalcService_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalcService_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _CalcService_Divide_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Average",
			Handler:       _CalcService_Average_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pb/calc.proto",
}
