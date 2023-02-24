package from

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/v3nooonn/trytry/pkg/constant/types"
	"github.com/v3nooonn/trytry/pkg/utils/contextx"
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
