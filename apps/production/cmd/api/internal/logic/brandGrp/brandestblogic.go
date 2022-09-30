package brandGrp

import (
	"context"
	"net/http"
	"time"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/model/brand"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandEstbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandEstbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandEstbLogic {
	return &BrandEstbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandEstbLogic) BrandEstb(req *types.BrandEstbReq) (resp *types.BrandEstbResp, err error) {
	now := time.Now().UTC()

	result, err := l.svcCtx.ProductionBrandModel.Insert(l.ctx, &brand.ProductionBrand{
		Category:  int64(req.Category),
		Name:      req.Brand,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	resp = &types.BrandEstbResp{
		Code:    http.StatusOK,
		Message: "ok",
		Data: types.Brand{
			ID:       uint64(id),
			Category: req.Category,
			Name:     req.Brand,
		},
	}
	return
}
