// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fullcore/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryAllObitRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllObitRequest) Reset()         { *m = QueryAllObitRequest{} }
func (m *QueryAllObitRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllObitRequest) ProtoMessage()    {}
func (*QueryAllObitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_90487403b6db27d2, []int{0}
}
func (m *QueryAllObitRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllObitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllObitRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllObitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllObitRequest.Merge(m, src)
}
func (m *QueryAllObitRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllObitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllObitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllObitRequest proto.InternalMessageInfo

func (m *QueryAllObitRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllObitResponse struct {
	Obit       []*Obit             `protobuf:"bytes,1,rep,name=Obit,proto3" json:"Obit,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllObitResponse) Reset()         { *m = QueryAllObitResponse{} }
func (m *QueryAllObitResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllObitResponse) ProtoMessage()    {}
func (*QueryAllObitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90487403b6db27d2, []int{1}
}
func (m *QueryAllObitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllObitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllObitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllObitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllObitResponse.Merge(m, src)
}
func (m *QueryAllObitResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllObitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllObitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllObitResponse proto.InternalMessageInfo

func (m *QueryAllObitResponse) GetObit() []*Obit {
	if m != nil {
		return m.Obit
	}
	return nil
}

func (m *QueryAllObitResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryGetObitRequest struct {
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
}

func (m *QueryGetObitRequest) Reset()         { *m = QueryGetObitRequest{} }
func (m *QueryGetObitRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetObitRequest) ProtoMessage()    {}
func (*QueryGetObitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_90487403b6db27d2, []int{2}
}
func (m *QueryGetObitRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetObitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetObitRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetObitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetObitRequest.Merge(m, src)
}
func (m *QueryGetObitRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetObitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetObitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetObitRequest proto.InternalMessageInfo

func (m *QueryGetObitRequest) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

type QueryGetObitResponse struct {
	Obit *Obit `protobuf:"bytes,1,opt,name=Obit,proto3" json:"Obit,omitempty"`
}

func (m *QueryGetObitResponse) Reset()         { *m = QueryGetObitResponse{} }
func (m *QueryGetObitResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetObitResponse) ProtoMessage()    {}
func (*QueryGetObitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90487403b6db27d2, []int{3}
}
func (m *QueryGetObitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetObitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetObitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetObitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetObitResponse.Merge(m, src)
}
func (m *QueryGetObitResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetObitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetObitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetObitResponse proto.InternalMessageInfo

func (m *QueryGetObitResponse) GetObit() *Obit {
	if m != nil {
		return m.Obit
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllObitRequest)(nil), "obadafoundation.fullcore.fullcore.QueryAllObitRequest")
	proto.RegisterType((*QueryAllObitResponse)(nil), "obadafoundation.fullcore.fullcore.QueryAllObitResponse")
	proto.RegisterType((*QueryGetObitRequest)(nil), "obadafoundation.fullcore.fullcore.QueryGetObitRequest")
	proto.RegisterType((*QueryGetObitResponse)(nil), "obadafoundation.fullcore.fullcore.QueryGetObitResponse")
}

func init() { proto.RegisterFile("fullcore/query.proto", fileDescriptor_90487403b6db27d2) }

var fileDescriptor_90487403b6db27d2 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2b, 0xcd, 0xc9,
	0x49, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0xcc, 0x4f, 0x4a, 0x4c, 0x49, 0x4c, 0xcb, 0x2f, 0xcd, 0x4b, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0xd3, 0x83, 0xa9, 0x82, 0x33, 0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13,
	0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0xc0, 0x0a, 0x8b, 0x21, 0x06, 0x48, 0x69, 0x25,
	0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xeb, 0x27, 0x25, 0x16, 0x43, 0x4d, 0xd6, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x48, 0x4c, 0xcf, 0xcc, 0x83, 0x98, 0x0a, 0x51, 0x2b, 0x0c, 0x77,
	0x42, 0x7e, 0x52, 0x66, 0x09, 0x44, 0x50, 0x29, 0x96, 0x4b, 0x38, 0x10, 0xa4, 0xcd, 0x31, 0x27,
	0xc7, 0x3f, 0x29, 0xb3, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc8, 0x8d, 0x8b, 0x0b,
	0xa1, 0x5f, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4d, 0x0f, 0x62, 0x99, 0x1e, 0xc8, 0x32,
	0x3d, 0x88, 0x37, 0xa0, 0x96, 0xe9, 0x05, 0x24, 0xa6, 0xa7, 0x42, 0xf5, 0x06, 0x21, 0xe9, 0x54,
	0x9a, 0xc3, 0xc8, 0x25, 0x82, 0x6a, 0x7e, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x35, 0x17,
	0x0b, 0x88, 0x2f, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xae, 0x47, 0x30, 0x20, 0xf4, 0xc0,
	0xda, 0xc1, 0x9a, 0x84, 0xdc, 0x51, 0x5c, 0xc7, 0x04, 0x76, 0x9d, 0x3a, 0x41, 0xd7, 0x41, 0x6c,
	0x46, 0x71, 0x9e, 0x3a, 0xd4, 0xf7, 0xee, 0xa9, 0x25, 0xc8, 0xbe, 0x17, 0xe0, 0x62, 0x4e, 0xc9,
	0x4c, 0x01, 0x7b, 0x9b, 0x33, 0x08, 0xc4, 0x54, 0x0a, 0x86, 0x7a, 0x03, 0xae, 0x10, 0xc3, 0x1b,
	0x8c, 0x24, 0x7b, 0xc3, 0xe8, 0x05, 0x13, 0x17, 0x2b, 0xd8, 0x54, 0xa1, 0xf5, 0x8c, 0x10, 0x73,
	0x84, 0xcc, 0x88, 0x30, 0x01, 0x8b, 0x8b, 0xa5, 0xcc, 0x49, 0xd6, 0x07, 0xf1, 0x80, 0x92, 0x59,
	0xd3, 0xe5, 0x27, 0x93, 0x99, 0x0c, 0x84, 0xf4, 0xf4, 0xc1, 0x06, 0xe8, 0x22, 0x4c, 0xd0, 0x87,
	0x27, 0x17, 0x94, 0x74, 0xa3, 0x5f, 0x9d, 0x92, 0x99, 0x52, 0x2b, 0xb4, 0x86, 0x91, 0x8b, 0x1d,
	0x64, 0x90, 0x63, 0x4e, 0x0e, 0xf1, 0x8e, 0x46, 0x4d, 0x64, 0xc4, 0x3b, 0x1a, 0x2d, 0xf1, 0x28,
	0x19, 0x80, 0x1d, 0xad, 0x25, 0xa4, 0x41, 0xac, 0xa3, 0x9d, 0xfc, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x17, 0x8f, 0x69, 0x15, 0x08, 0x66, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xfb, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0xbe, 0x8c, 0x6d, 0xd8, 0x03, 0x00, 0x00,
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
	Obit(ctx context.Context, in *QueryGetObitRequest, opts ...grpc.CallOption) (*QueryGetObitResponse, error)
	ObitAll(ctx context.Context, in *QueryAllObitRequest, opts ...grpc.CallOption) (*QueryAllObitResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Obit(ctx context.Context, in *QueryGetObitRequest, opts ...grpc.CallOption) (*QueryGetObitResponse, error) {
	out := new(QueryGetObitResponse)
	err := c.cc.Invoke(ctx, "/obadafoundation.fullcore.fullcore.Query/Obit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ObitAll(ctx context.Context, in *QueryAllObitRequest, opts ...grpc.CallOption) (*QueryAllObitResponse, error) {
	out := new(QueryAllObitResponse)
	err := c.cc.Invoke(ctx, "/obadafoundation.fullcore.fullcore.Query/ObitAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Obit(context.Context, *QueryGetObitRequest) (*QueryGetObitResponse, error)
	ObitAll(context.Context, *QueryAllObitRequest) (*QueryAllObitResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Obit(ctx context.Context, req *QueryGetObitRequest) (*QueryGetObitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Obit not implemented")
}
func (*UnimplementedQueryServer) ObitAll(ctx context.Context, req *QueryAllObitRequest) (*QueryAllObitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObitAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Obit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetObitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Obit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/obadafoundation.fullcore.fullcore.Query/Obit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Obit(ctx, req.(*QueryGetObitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ObitAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllObitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ObitAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/obadafoundation.fullcore.fullcore.Query/ObitAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ObitAll(ctx, req.(*QueryAllObitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "obadafoundation.fullcore.fullcore.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Obit",
			Handler:    _Query_Obit_Handler,
		},
		{
			MethodName: "ObitAll",
			Handler:    _Query_ObitAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fullcore/query.proto",
}

func (m *QueryAllObitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllObitRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllObitRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllObitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllObitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllObitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Obit) > 0 {
		for iNdEx := len(m.Obit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Obit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGetObitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetObitRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetObitRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetObitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetObitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetObitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Obit != nil {
		{
			size, err := m.Obit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
func (m *QueryAllObitRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllObitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Obit) > 0 {
		for _, e := range m.Obit {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetObitRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetObitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Obit != nil {
		l = m.Obit.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllObitRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllObitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllObitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllObitResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllObitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllObitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Obit", wireType)
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
			m.Obit = append(m.Obit, &Obit{})
			if err := m.Obit[len(m.Obit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGetObitRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetObitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetObitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
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
			m.Did = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetObitResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetObitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetObitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Obit", wireType)
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
			if m.Obit == nil {
				m.Obit = &Obit{}
			}
			if err := m.Obit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
