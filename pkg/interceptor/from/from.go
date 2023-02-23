package from

import (
	"context"
	"fmt"

	"github.com/v3nooonn/trytry/pkg/constant/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TenantInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// TODO: return error
		fmt.Println("error")
	}

	s := md.Get(types.ContextKey("key").Key())
	if len(s) == 0 {
		// TODO: return error
		fmt.Println("error")
	}

	return handler(context.WithValue(ctx, types.ContextKey("k"), "v"), req)
}
