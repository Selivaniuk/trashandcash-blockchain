// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trashandcashblockchain/trashandcashblockchain/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbcd9dae7bfaea6c, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbcd9dae7bfaea6c, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetSpinRequest struct {
	//id is a unique identifier of the Spin
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetSpinRequest) Reset()         { *m = QueryGetSpinRequest{} }
func (m *QueryGetSpinRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSpinRequest) ProtoMessage()    {}
func (*QueryGetSpinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbcd9dae7bfaea6c, []int{2}
}
func (m *QueryGetSpinRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSpinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSpinRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSpinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSpinRequest.Merge(m, src)
}
func (m *QueryGetSpinRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSpinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSpinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSpinRequest proto.InternalMessageInfo

func (m *QueryGetSpinRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetSpinResponse struct {
	Spin Spin `protobuf:"bytes,1,opt,name=Spin,proto3" json:"Spin"`
}

func (m *QueryGetSpinResponse) Reset()         { *m = QueryGetSpinResponse{} }
func (m *QueryGetSpinResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetSpinResponse) ProtoMessage()    {}
func (*QueryGetSpinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbcd9dae7bfaea6c, []int{3}
}
func (m *QueryGetSpinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSpinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSpinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSpinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSpinResponse.Merge(m, src)
}
func (m *QueryGetSpinResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSpinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSpinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSpinResponse proto.InternalMessageInfo

func (m *QueryGetSpinResponse) GetSpin() Spin {
	if m != nil {
		return m.Spin
	}
	return Spin{}
}

type QuerySpinsRequest struct {
	//address is the account address string.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QuerySpinsRequest) Reset()         { *m = QuerySpinsRequest{} }
func (m *QuerySpinsRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySpinsRequest) ProtoMessage()    {}
func (*QuerySpinsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbcd9dae7bfaea6c, []int{4}
}
func (m *QuerySpinsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpinsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpinsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpinsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpinsRequest.Merge(m, src)
}
func (m *QuerySpinsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpinsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpinsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpinsRequest proto.InternalMessageInfo

func (m *QuerySpinsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QuerySpinsResponse struct {
	Spin []Spin `protobuf:"bytes,1,rep,name=Spin,proto3" json:"Spin"`
}

func (m *QuerySpinsResponse) Reset()         { *m = QuerySpinsResponse{} }
func (m *QuerySpinsResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySpinsResponse) ProtoMessage()    {}
func (*QuerySpinsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbcd9dae7bfaea6c, []int{5}
}
func (m *QuerySpinsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpinsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpinsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpinsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpinsResponse.Merge(m, src)
}
func (m *QuerySpinsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpinsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpinsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpinsResponse proto.InternalMessageInfo

func (m *QuerySpinsResponse) GetSpin() []Spin {
	if m != nil {
		return m.Spin
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "trashandcashblockchain.trashandcashblockchain.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "trashandcashblockchain.trashandcashblockchain.QueryParamsResponse")
	proto.RegisterType((*QueryGetSpinRequest)(nil), "trashandcashblockchain.trashandcashblockchain.QueryGetSpinRequest")
	proto.RegisterType((*QueryGetSpinResponse)(nil), "trashandcashblockchain.trashandcashblockchain.QueryGetSpinResponse")
	proto.RegisterType((*QuerySpinsRequest)(nil), "trashandcashblockchain.trashandcashblockchain.QuerySpinsRequest")
	proto.RegisterType((*QuerySpinsResponse)(nil), "trashandcashblockchain.trashandcashblockchain.QuerySpinsResponse")
}

func init() {
	proto.RegisterFile("trashandcashblockchain/trashandcashblockchain/query.proto", fileDescriptor_dbcd9dae7bfaea6c)
}

var fileDescriptor_dbcd9dae7bfaea6c = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xcf, 0xac, 0xe9, 0x8a, 0x23, 0x08, 0x8e, 0x7b, 0x28, 0x41, 0xa2, 0x04, 0x04, 0x11, 0x92,
	0x61, 0xbb, 0xb4, 0xb4, 0x1e, 0xc4, 0xae, 0x07, 0x4f, 0x42, 0x4d, 0x6f, 0xde, 0x26, 0xc9, 0x90,
	0x1d, 0xed, 0xce, 0x4c, 0x33, 0x53, 0xd9, 0x22, 0x5e, 0xfc, 0x04, 0x82, 0x1f, 0xc9, 0x4b, 0xf1,
	0x62, 0xc1, 0x8b, 0x27, 0x91, 0x5d, 0x8f, 0x7e, 0x08, 0xc9, 0xcc, 0x04, 0x1a, 0xdb, 0x05, 0xb3,
	0xeb, 0x2d, 0xef, 0xe5, 0xbd, 0xdf, 0x9f, 0xf7, 0x5e, 0x02, 0xf7, 0x74, 0x45, 0xd4, 0x84, 0xf0,
	0x22, 0x27, 0x6a, 0x92, 0x1d, 0x89, 0xfc, 0x4d, 0x3e, 0x21, 0x8c, 0xe3, 0x25, 0xe9, 0xe3, 0x13,
	0x5a, 0x9d, 0x26, 0xb2, 0x12, 0x5a, 0xa0, 0xf8, 0xea, 0x9a, 0xe4, 0xea, 0x74, 0x30, 0x28, 0x45,
	0x29, 0x4c, 0x27, 0xae, 0x9f, 0x2c, 0x48, 0x70, 0xb7, 0x14, 0xa2, 0x3c, 0xa2, 0x98, 0x48, 0x86,
	0x09, 0xe7, 0x42, 0x13, 0xcd, 0x04, 0x57, 0xee, 0xed, 0xa3, 0x5c, 0xa8, 0xa9, 0x50, 0x38, 0x23,
	0x8a, 0x5a, 0x6e, 0xfc, 0x76, 0x98, 0x51, 0x4d, 0x86, 0x58, 0x92, 0x92, 0x71, 0x53, 0xec, 0x6a,
	0x1f, 0x77, 0x73, 0x22, 0x49, 0x45, 0xa6, 0x0d, 0xcf, 0x6e, 0xb7, 0x5e, 0x25, 0x99, 0x63, 0x8d,
	0x06, 0x10, 0xbd, 0xac, 0x75, 0x1d, 0x18, 0xb8, 0x94, 0x1e, 0x9f, 0x50, 0xa5, 0xa3, 0xd7, 0xf0,
	0x4e, 0x2b, 0xab, 0xa4, 0xe0, 0x8a, 0xa2, 0x43, 0xd8, 0xb7, 0xb4, 0x9b, 0xe0, 0x3e, 0x78, 0x78,
	0x73, 0x6b, 0x3b, 0xe9, 0x34, 0xc2, 0xc4, 0xc2, 0x8d, 0xfd, 0xb3, 0x1f, 0xf7, 0xbc, 0xd4, 0x41,
	0x45, 0x0f, 0x1c, 0xd7, 0x73, 0xaa, 0x0f, 0x25, 0xe3, 0x4e, 0x02, 0xba, 0x05, 0x7b, 0xac, 0x30,
	0x3c, 0x7e, 0xda, 0x63, 0x45, 0x44, 0xe1, 0xa0, 0x5d, 0xe6, 0x34, 0xbd, 0x80, 0x7e, 0x1d, 0x3b,
	0x45, 0xa3, 0x8e, 0x8a, 0xea, 0x56, 0xa7, 0xc7, 0xc0, 0x44, 0x31, 0xbc, 0x6d, 0x68, 0xea, 0xa0,
	0x19, 0x07, 0xda, 0x84, 0xd7, 0x49, 0x51, 0x54, 0x54, 0x59, 0xe3, 0x37, 0xd2, 0x26, 0x8c, 0x72,
	0x37, 0x3e, 0x57, 0x7e, 0x49, 0xd3, 0xb5, 0xff, 0xa0, 0x69, 0xeb, 0xb7, 0x0f, 0x37, 0x0c, 0x0b,
	0xfa, 0x0a, 0x60, 0xdf, 0x0e, 0x11, 0xed, 0x77, 0x44, 0xbd, 0xbc, 0xe5, 0x60, 0xbc, 0x0e, 0x84,
	0xb5, 0x1a, 0x3d, 0xf9, 0xf0, 0xed, 0xd7, 0xa7, 0xde, 0x2e, 0xda, 0x69, 0xdd, 0x5a, 0xfc, 0xaf,
	0xf7, 0x8b, 0xbe, 0x00, 0x3b, 0x2b, 0xb4, 0x92, 0x98, 0xf6, 0xcd, 0x04, 0xcf, 0xd6, 0xc2, 0x70,
	0x8e, 0xf6, 0x8c, 0xa3, 0x11, 0x1a, 0x62, 0x32, 0x63, 0x5c, 0x4c, 0xd9, 0xec, 0xa2, 0x9b, 0x26,
	0xf7, 0xd7, 0xd7, 0x84, 0xdf, 0xb1, 0xe2, 0x3d, 0xfa, 0x0c, 0xe0, 0x86, 0xb9, 0x04, 0xf4, 0x74,
	0x15, 0x25, 0x17, 0x6f, 0x2e, 0xd8, 0x5f, 0x03, 0xc1, 0x39, 0xd9, 0x36, 0x4e, 0x30, 0x8a, 0xbb,
	0x38, 0x51, 0xe3, 0x83, 0xb3, 0x79, 0x08, 0xce, 0xe7, 0x21, 0xf8, 0x39, 0x0f, 0xc1, 0xc7, 0x45,
	0xe8, 0x9d, 0x2f, 0x42, 0xef, 0xfb, 0x22, 0xf4, 0x5e, 0xed, 0x2c, 0xdb, 0xf1, 0x6c, 0xd9, 0x96,
	0xf5, 0xa9, 0xa4, 0x2a, 0xeb, 0x9b, 0x7f, 0xcd, 0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0x76, 0xcb, 0x46, 0xad, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries information about a Spin by its ID.
	Spin(ctx context.Context, in *QueryGetSpinRequest, opts ...grpc.CallOption) (*QueryGetSpinResponse, error)
	// Queries a list of Spins.
	Spins(ctx context.Context, in *QuerySpinsRequest, opts ...grpc.CallOption) (*QuerySpinsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/trashandcashblockchain.trashandcashblockchain.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Spin(ctx context.Context, in *QueryGetSpinRequest, opts ...grpc.CallOption) (*QueryGetSpinResponse, error) {
	out := new(QueryGetSpinResponse)
	err := c.cc.Invoke(ctx, "/trashandcashblockchain.trashandcashblockchain.Query/Spin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Spins(ctx context.Context, in *QuerySpinsRequest, opts ...grpc.CallOption) (*QuerySpinsResponse, error) {
	out := new(QuerySpinsResponse)
	err := c.cc.Invoke(ctx, "/trashandcashblockchain.trashandcashblockchain.Query/Spins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries information about a Spin by its ID.
	Spin(context.Context, *QueryGetSpinRequest) (*QueryGetSpinResponse, error)
	// Queries a list of Spins.
	Spins(context.Context, *QuerySpinsRequest) (*QuerySpinsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Spin(ctx context.Context, req *QueryGetSpinRequest) (*QueryGetSpinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spin not implemented")
}
func (*UnimplementedQueryServer) Spins(ctx context.Context, req *QuerySpinsRequest) (*QuerySpinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spins not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trashandcashblockchain.trashandcashblockchain.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Spin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSpinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Spin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trashandcashblockchain.trashandcashblockchain.Query/Spin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Spin(ctx, req.(*QueryGetSpinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Spins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Spins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trashandcashblockchain.trashandcashblockchain.Query/Spins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Spins(ctx, req.(*QuerySpinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trashandcashblockchain.trashandcashblockchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Spin",
			Handler:    _Query_Spin_Handler,
		},
		{
			MethodName: "Spins",
			Handler:    _Query_Spins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trashandcashblockchain/trashandcashblockchain/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetSpinRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSpinRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSpinRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetSpinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSpinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSpinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySpinsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpinsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpinsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySpinsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpinsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpinsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Spin) > 0 {
		for iNdEx := len(m.Spin) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Spin[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetSpinRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetSpinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Spin.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySpinsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySpinsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Spin) > 0 {
		for _, e := range m.Spin {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSpinRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSpinRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSpinRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSpinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSpinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSpinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySpinsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySpinsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpinsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySpinsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySpinsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpinsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spin = append(m.Spin, Spin{})
			if err := m.Spin[len(m.Spin)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
