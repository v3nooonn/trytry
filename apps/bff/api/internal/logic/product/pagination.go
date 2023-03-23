package product

import (
	"context"
	"fmt"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	"github.com/v3nooonn/trytry/apps/brand/brandclient"
	"github.com/v3nooonn/trytry/apps/category/categoryclient"
	"github.com/v3nooonn/trytry/apps/product/productclient"
	e "github.com/v3nooonn/trytry/pkg/expands/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
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
	psResp, err := l.svcCtx.ProductionRPC.Pagination(l.ctx, &productclient.PaginationReq{
		Cursor: req.Cursor,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	iProducts, err := mr.MapReduce(
		func(source chan<- interface{}) {
			for _, p := range psResp.GetProducts() {
				source <- p
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			//p, ok := item.(types.Production)
			p, ok := item.(*productclient.ProductInfo)
			if !ok {
				cancel(e.Internal(fmt.Sprintf("casting to %T occurred an error", types.Production{})))
			}

			brand, err := l.svcCtx.BrandRPC.Info(l.ctx, &brandclient.InfoReq{BrandId: p.BrandId})
			if err != nil {
				cancel(err)
				return
			}
			category, err := l.svcCtx.CategoryRPC.Info(l.ctx, &categoryclient.InfoReq{CategoryId: p.CategoryId})
			if err != nil {
				cancel(err)
				return
			}

			writer.Write(types.Production{
				ID:           p.GetId(),
				Category:     category.GetName(),
				Brand:        brand.GetName(),
				Series:       p.GetSeries(),
				Name:         p.GetName(),
				Abbreviation: p.GetAbbreviation(),
				CreatedAt:    p.GetCreatedAt(),
				UpdatedAt:    p.GetUpdatedAt(),
			})
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			prods := make([]types.Production, 0, len(psResp.GetProducts()))
			for p := range pipe {
				p, ok := p.(types.Production)
				if !ok {
					cancel(e.Internal(fmt.Sprintf("casting to %T occurred an error", types.Production{})))
				}

				prods = append(prods, p)
			}

			writer.Write(prods)
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.PageResp{
		Pagination: types.PgnResp{
			Prev: "prev",
			Next: "next",
		},
		Productions: iProducts.([]types.Production),
	}, nil
}
