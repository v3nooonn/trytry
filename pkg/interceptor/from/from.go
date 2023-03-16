package from

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/v3nooonn/trytry/pkg/constant/types"
	"github.com/v3nooonn/trytry/pkg/expands/contextx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Tenant(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.Wrap(fmt.Errorf("missing metadata"), "missing metadata")
	}

	s := md.Get(types.CtxKeyTenantSchema.Key())
	if len(s) == 0 {
		return nil, errors.Wrap(fmt.Errorf("empty metadata"), "empty metadata")
	}

	return handler(
		contextx.WithContext(ctx).
			WithValue(types.CtxKeyTenantSchema, s).Context(),
		req,
	)
}

//func TempFrom(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return errors.Wrap(fmt.Errorf("missing metadata"), "missing metadata")
//	}
//
//	s := md.Get(types.CtxKeyTenantSchema.Key())
//	if len(s) == 0 {
//		return errors.Wrap(fmt.Errorf("empty metadata"), "empty metadata")
//	}
//
//	return invoker(
//		metadata.NewOutgoingContext(
//			ctx,
//			md,
//		),
//		method, req, reply, cc, opts...,
//	)
//}
