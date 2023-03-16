package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// ErrLogInterceptor log
func ErrLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		// write log to...
		fmt.Printf("Error Log interceptro: %s\n", err.Error())
		return resp, err
	}

	return resp, err
}
