package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/v3nooonn/trytry/apps/production/internal/svc"
	"github.com/v3nooonn/trytry/apps/production/pb/production"
)

type ViewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewsLogic {
	return &ViewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewsLogic) Views(in *production.ViewsReq) (*production.ViewsResp, error) {
	listResp, err := l.svcCtx.Production.List(l.ctx)
	if err != nil {
		return nil, err
	}

	pvs := make([]*production.ProdView, 0, len(listResp))
	for _, prod := range listResp {
		pvs = append(pvs, &production.ProdView{
			Id:        prod.Id,
			Name:      prod.Name,
			Category:  prod.Category,
			CreatedAt: prod.CreatedAt,
			UpdatedAt: prod.UpdatedAt,
		})
	}

	return &production.ViewsResp{ProdViews: pvs}, nil
}
