package product

import (
	"context"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	"github.com/v3nooonn/trytry/apps/product/pb/product"
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

func (l *PaginationLogic) Pagination(req *types.PgnReq) (*types.PageResp, error) {
	rpcResp, err := l.svcCtx.ProductionRPC.Pagination(l.ctx, &product.PaginationReq{
		Cursor: req.Cursor,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	prods := make([]types.Production, 0, len(rpcResp.GetProducts()))
	for _, prod := range rpcResp.GetProducts() {
		prods = append(prods, types.Production{
			ID:           prod.GetId(),
			Brand:        "prod.GetBrandId()",
			Category:     "prod.GetCategoryId()",
			Series:       prod.GetSeries(),
			Name:         prod.GetName(),
			Abbreviation: prod.GetAbbreviation(),
			CreatedAt:    prod.GetCreatedAt(),
			UpdatedAt:    prod.GetUpdatedAt(),
		})
	}

	return &types.PageResp{
		Pagination: types.PgnResp{
			Prev: "prev",
			Next: "next",
		},
		Productions: prods,
	}, nil
}
