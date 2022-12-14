// Code generated by goctl. DO NOT EDIT!
// Source: customer.proto

package customerclient

import (
	"context"

	"github.com/v3nooonn/trytry/apps/customer/cmd/rpc/pb/customer"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoByEmailReq = customer.InfoByEmailReq
	InfoReq        = customer.InfoReq
	InfoResp       = customer.InfoResp

	Customer interface {
		Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		InfoByEmail(ctx context.Context, in *InfoByEmailReq, opts ...grpc.CallOption) (*InfoResp, error)
	}

	defaultCustomer struct {
		cli zrpc.Client
	}
)

func NewCustomer(cli zrpc.Client) Customer {
	return &defaultCustomer{
		cli: cli,
	}
}

func (m *defaultCustomer) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := customer.NewCustomerClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}

func (m *defaultCustomer) InfoByEmail(ctx context.Context, in *InfoByEmailReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := customer.NewCustomerClient(m.cli.Conn())
	return client.InfoByEmail(ctx, in, opts...)
}
