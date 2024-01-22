// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: obadafoundation/obit/v1/tx.proto

package obitv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateNFT_FullMethodName        = "/obadafoundation.obit.v1.Msg/UpdateNFT"
	Msg_UpdateUriHash_FullMethodName    = "/obadafoundation.obit.v1.Msg/UpdateUriHash"
	Msg_MintNFT_FullMethodName          = "/obadafoundation.obit.v1.Msg/MintNFT"
	Msg_BatchMintNFT_FullMethodName     = "/obadafoundation.obit.v1.Msg/BatchMintNFT"
	Msg_TransferNFT_FullMethodName      = "/obadafoundation.obit.v1.Msg/TransferNFT"
	Msg_BatchTransferNFT_FullMethodName = "/obadafoundation.obit.v1.Msg/BatchTransferNFT"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateNFT updates NFT properties (definition will come in future)
	UpdateNFT(ctx context.Context, in *MsgUpdateNFT, opts ...grpc.CallOption) (*MsgUpdateNFTResponse, error)
	// UpdateUriHash updates URI hash for existing NFT
	UpdateUriHash(ctx context.Context, in *MsgUpdateUriHash, opts ...grpc.CallOption) (*MsgUpdateUriHashResponse, error)
	// MintNFT mint a new OBADA NFT (Obit)
	MintNFT(ctx context.Context, in *MsgMintNFT, opts ...grpc.CallOption) (*MsgMintNFTResponse, error)
	// BatchMintNFT mint many NFTs
	BatchMintNFT(ctx context.Context, in *MsgBatchMintNFT, opts ...grpc.CallOption) (*MsgBatchMintNFTResponse, error)
	// TransferNFT send NFT to the new owner address
	TransferNFT(ctx context.Context, in *MsgTransferNFT, opts ...grpc.CallOption) (*MsgTransferNFTResponse, error)
	// BatchTransferNFT transfer a batch NFTs
	BatchTransferNFT(ctx context.Context, in *MsgBatchTransferNFT, opts ...grpc.CallOption) (*MsgBatchTransferNFTResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateNFT(ctx context.Context, in *MsgUpdateNFT, opts ...grpc.CallOption) (*MsgUpdateNFTResponse, error) {
	out := new(MsgUpdateNFTResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateNFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateUriHash(ctx context.Context, in *MsgUpdateUriHash, opts ...grpc.CallOption) (*MsgUpdateUriHashResponse, error) {
	out := new(MsgUpdateUriHashResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateUriHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintNFT(ctx context.Context, in *MsgMintNFT, opts ...grpc.CallOption) (*MsgMintNFTResponse, error) {
	out := new(MsgMintNFTResponse)
	err := c.cc.Invoke(ctx, Msg_MintNFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BatchMintNFT(ctx context.Context, in *MsgBatchMintNFT, opts ...grpc.CallOption) (*MsgBatchMintNFTResponse, error) {
	out := new(MsgBatchMintNFTResponse)
	err := c.cc.Invoke(ctx, Msg_BatchMintNFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferNFT(ctx context.Context, in *MsgTransferNFT, opts ...grpc.CallOption) (*MsgTransferNFTResponse, error) {
	out := new(MsgTransferNFTResponse)
	err := c.cc.Invoke(ctx, Msg_TransferNFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BatchTransferNFT(ctx context.Context, in *MsgBatchTransferNFT, opts ...grpc.CallOption) (*MsgBatchTransferNFTResponse, error) {
	out := new(MsgBatchTransferNFTResponse)
	err := c.cc.Invoke(ctx, Msg_BatchTransferNFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateNFT updates NFT properties (definition will come in future)
	UpdateNFT(context.Context, *MsgUpdateNFT) (*MsgUpdateNFTResponse, error)
	// UpdateUriHash updates URI hash for existing NFT
	UpdateUriHash(context.Context, *MsgUpdateUriHash) (*MsgUpdateUriHashResponse, error)
	// MintNFT mint a new OBADA NFT (Obit)
	MintNFT(context.Context, *MsgMintNFT) (*MsgMintNFTResponse, error)
	// BatchMintNFT mint many NFTs
	BatchMintNFT(context.Context, *MsgBatchMintNFT) (*MsgBatchMintNFTResponse, error)
	// TransferNFT send NFT to the new owner address
	TransferNFT(context.Context, *MsgTransferNFT) (*MsgTransferNFTResponse, error)
	// BatchTransferNFT transfer a batch NFTs
	BatchTransferNFT(context.Context, *MsgBatchTransferNFT) (*MsgBatchTransferNFTResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateNFT(context.Context, *MsgUpdateNFT) (*MsgUpdateNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNFT not implemented")
}
func (UnimplementedMsgServer) UpdateUriHash(context.Context, *MsgUpdateUriHash) (*MsgUpdateUriHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUriHash not implemented")
}
func (UnimplementedMsgServer) MintNFT(context.Context, *MsgMintNFT) (*MsgMintNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintNFT not implemented")
}
func (UnimplementedMsgServer) BatchMintNFT(context.Context, *MsgBatchMintNFT) (*MsgBatchMintNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchMintNFT not implemented")
}
func (UnimplementedMsgServer) TransferNFT(context.Context, *MsgTransferNFT) (*MsgTransferNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferNFT not implemented")
}
func (UnimplementedMsgServer) BatchTransferNFT(context.Context, *MsgBatchTransferNFT) (*MsgBatchTransferNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchTransferNFT not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateNFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateNFT(ctx, req.(*MsgUpdateNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateUriHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateUriHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateUriHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateUriHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateUriHash(ctx, req.(*MsgUpdateUriHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintNFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintNFT(ctx, req.(*MsgMintNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BatchMintNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBatchMintNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BatchMintNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BatchMintNFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BatchMintNFT(ctx, req.(*MsgBatchMintNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferNFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferNFT(ctx, req.(*MsgTransferNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BatchTransferNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBatchTransferNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BatchTransferNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BatchTransferNFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BatchTransferNFT(ctx, req.(*MsgBatchTransferNFT))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "obadafoundation.obit.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateNFT",
			Handler:    _Msg_UpdateNFT_Handler,
		},
		{
			MethodName: "UpdateUriHash",
			Handler:    _Msg_UpdateUriHash_Handler,
		},
		{
			MethodName: "MintNFT",
			Handler:    _Msg_MintNFT_Handler,
		},
		{
			MethodName: "BatchMintNFT",
			Handler:    _Msg_BatchMintNFT_Handler,
		},
		{
			MethodName: "TransferNFT",
			Handler:    _Msg_TransferNFT_Handler,
		},
		{
			MethodName: "BatchTransferNFT",
			Handler:    _Msg_BatchTransferNFT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "obadafoundation/obit/v1/tx.proto",
}
