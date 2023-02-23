package outgoing

import (
	"context"

	"github.com/v3nooonn/trytry/pkg/constant/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TenantInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := ctx.Value(types.ContextKey("key")).(metadata.MD)

	return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...)
}
