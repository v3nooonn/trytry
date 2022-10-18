package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// interceptors
func ErrLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		// write log to...
		return resp, err
	} else {
		return resp, err
	}
}
