package product

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	"github.com/v3nooonn/trytry/apps/product/pb/product"
	e "github.com/v3nooonn/trytry/pkg/expands/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaginationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginationLogic {
	return &PaginationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaginationLogic) Pagination(req *types.PageReq) (*types.PageResp, error) {
	resp, err := l.svcCtx.ProductionRPC.Pagination(l.ctx, &product.PaginationReq{
		Cursor: req.Cursor,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	prods := make([]types.Production, 0, len(resp.GetProducts()))
	if err := copier.Copy(&prods, resp.GetProducts()); err != nil {
		return nil, e.Internal(err.Error())
	}

	return &types.PageResp{Productions: prods}, nil
}
