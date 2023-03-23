package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/v3nooonn/trytry/apps/product/internal/svc"
	"github.com/v3nooonn/trytry/apps/product/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaginationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginationLogic {
	return &PaginationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaginationLogic) Pagination(in *product.PaginationReq) (*product.PaginationResp, error) {
	prodsResp, err := l.svcCtx.Product.Pagination(l.ctx, in.Cursor, in.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "product pagination error")
	}

	products := make([]*product.ProductInfo, 0, len(prodsResp))
	if err := copier.Copy(&products, prodsResp); err != nil {
		return nil, errors.Wrap(err, "copier error")
	}

	return &product.PaginationResp{Products: products}, nil
}
