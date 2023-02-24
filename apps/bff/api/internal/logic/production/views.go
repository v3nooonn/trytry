package production

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	"github.com/v3nooonn/trytry/apps/production/productionservice"
	e "github.com/v3nooonn/trytry/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewsLogic {
	return &ViewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewsLogic) Views() (*types.Brands, error) {
	fmt.Println("cao")
	brandsResp, err := l.svcCtx.ProductionService.Views(l.ctx, &productionservice.ViewsReq{})
	if err != nil {
		return nil, err
	}

	brands := make([]types.ProdBrand, 0, len(brandsResp.GetProdViews()))

	if err := copier.Copy(&brands, brandsResp.GetProdViews()); err != nil {
		return nil, errors.Wrap(e.Internal(err.Error()), "")
	}

	return &types.Brands{Brands: brands}, nil
}
