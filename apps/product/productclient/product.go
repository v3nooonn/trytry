// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productclient

import (
	"context"

	"github.com/v3nooonn/trytry/apps/product/pb/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq        = product.InfoReq
	PaginationReq  = product.PaginationReq
	PaginationResp = product.PaginationResp
	ProductInfo    = product.ProductInfo

	Product interface {
		Pagination(ctx context.Context, in *PaginationReq, opts ...grpc.CallOption) (*PaginationResp, error)
		Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*ProductInfo, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Pagination(ctx context.Context, in *PaginationReq, opts ...grpc.CallOption) (*PaginationResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Pagination(ctx, in, opts...)
}

func (m *defaultProduct) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}
