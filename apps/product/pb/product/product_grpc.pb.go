// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package product

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

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	Pagination(ctx context.Context, in *PaginationReq, opts ...grpc.CallOption) (*PaginationResp, error)
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*ProductInfo, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) Pagination(ctx context.Context, in *PaginationReq, opts ...grpc.CallOption) (*PaginationResp, error) {
	out := new(PaginationResp)
	err := c.cc.Invoke(ctx, "/product.Product/Pagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/product.Product/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	Pagination(context.Context, *PaginationReq) (*PaginationResp, error)
	Info(context.Context, *InfoReq) (*ProductInfo, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) Pagination(context.Context, *PaginationReq) (*PaginationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pagination not implemented")
}
func (UnimplementedProductServer) Info(context.Context, *InfoReq) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_Pagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Pagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Pagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Pagination(ctx, req.(*PaginationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pagination",
			Handler:    _Product_Pagination_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Product_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/product.proto",
}