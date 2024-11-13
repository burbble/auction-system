// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: bid.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BidService_PlaceBid_FullMethodName = "/auction.BidService/PlaceBid"
	BidService_GetBid_FullMethodName   = "/auction.BidService/GetBid"
	BidService_ListBids_FullMethodName = "/auction.BidService/ListBids"
)

// BidServiceClient is the client API for BidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidServiceClient interface {
	PlaceBid(ctx context.Context, in *PlaceBidRequest, opts ...grpc.CallOption) (*PlaceBidResponse, error)
	GetBid(ctx context.Context, in *GetBidRequest, opts ...grpc.CallOption) (*GetBidResponse, error)
	ListBids(ctx context.Context, in *ListBidsRequest, opts ...grpc.CallOption) (*ListBidsResponse, error)
}

type bidServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBidServiceClient(cc grpc.ClientConnInterface) BidServiceClient {
	return &bidServiceClient{cc}
}

func (c *bidServiceClient) PlaceBid(ctx context.Context, in *PlaceBidRequest, opts ...grpc.CallOption) (*PlaceBidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceBidResponse)
	err := c.cc.Invoke(ctx, BidService_PlaceBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidServiceClient) GetBid(ctx context.Context, in *GetBidRequest, opts ...grpc.CallOption) (*GetBidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBidResponse)
	err := c.cc.Invoke(ctx, BidService_GetBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidServiceClient) ListBids(ctx context.Context, in *ListBidsRequest, opts ...grpc.CallOption) (*ListBidsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBidsResponse)
	err := c.cc.Invoke(ctx, BidService_ListBids_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BidServiceServer is the server API for BidService service.
// All implementations must embed UnimplementedBidServiceServer
// for forward compatibility.
type BidServiceServer interface {
	PlaceBid(context.Context, *PlaceBidRequest) (*PlaceBidResponse, error)
	GetBid(context.Context, *GetBidRequest) (*GetBidResponse, error)
	ListBids(context.Context, *ListBidsRequest) (*ListBidsResponse, error)
	mustEmbedUnimplementedBidServiceServer()
}

// UnimplementedBidServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBidServiceServer struct{}

func (UnimplementedBidServiceServer) PlaceBid(context.Context, *PlaceBidRequest) (*PlaceBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}
func (UnimplementedBidServiceServer) GetBid(context.Context, *GetBidRequest) (*GetBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBid not implemented")
}
func (UnimplementedBidServiceServer) ListBids(context.Context, *ListBidsRequest) (*ListBidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBids not implemented")
}
func (UnimplementedBidServiceServer) mustEmbedUnimplementedBidServiceServer() {}
func (UnimplementedBidServiceServer) testEmbeddedByValue()                    {}

// UnsafeBidServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidServiceServer will
// result in compilation errors.
type UnsafeBidServiceServer interface {
	mustEmbedUnimplementedBidServiceServer()
}

func RegisterBidServiceServer(s grpc.ServiceRegistrar, srv BidServiceServer) {
	// If the following call pancis, it indicates UnimplementedBidServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BidService_ServiceDesc, srv)
}

func _BidService_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidService_PlaceBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).PlaceBid(ctx, req.(*PlaceBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidService_GetBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).GetBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidService_GetBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).GetBid(ctx, req.(*GetBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidService_ListBids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).ListBids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidService_ListBids_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).ListBids(ctx, req.(*ListBidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BidService_ServiceDesc is the grpc.ServiceDesc for BidService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auction.BidService",
	HandlerType: (*BidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBid",
			Handler:    _BidService_PlaceBid_Handler,
		},
		{
			MethodName: "GetBid",
			Handler:    _BidService_GetBid_Handler,
		},
		{
			MethodName: "ListBids",
			Handler:    _BidService_ListBids_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bid.proto",
}