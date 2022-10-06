package carGrp

import (
	"context"
	"net/http"
	"time"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/model/car"

	"github.com/zeromicro/go-zero/core/logx"
)

type EstbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEstbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EstbLogic {
	return &EstbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EstbLogic) Estb(req *types.CarEstbReq) (resp *types.CarEstbResp, err error) {
	now := time.Now().UTC()

	result, err := l.svcCtx.ProductionCarModel.Insert(l.ctx, &car.ProductionCar{
		Brand:     req.Brand,
		Serie:     req.Serie,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	resp = &types.CarEstbResp{
		Code:    http.StatusOK,
		Message: "ok",
		Data: types.Car{
			ID:    uint64(id),
			Brand: req.Brand,
			Serie: req.Serie,
		},
	}
	return
}
