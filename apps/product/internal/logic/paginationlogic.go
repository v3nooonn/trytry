package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &product.PaginationResp{}, nil
}
