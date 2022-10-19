package api

import (
	"context"

	"github.com/v3nooonn/trytry/common/constant"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TenantInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := ctx.Value(constant.ContKey("key")).(metadata.MD)

	return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...)
}
