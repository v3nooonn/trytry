package outgoing

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Tenant(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(
		metadata.NewOutgoingContext(
			ctx,
			NewMDGenerator(ctx).Metadata(),
		),
		method, req, reply, cc, opts...,
	)
}
