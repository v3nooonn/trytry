package logic

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/v3nooonn/trytry/apps/production/internal/svc"
	"github.com/v3nooonn/trytry/apps/production/pb/production"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCarInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarInfoLogic {
	return &CarInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CarInfoLogic) CarInfo(in *production.CarInfoReq) (*production.CarInfoResp, error) {
	car, err := l.svcCtx.ProductionCar.FindOne(l.ctx, in.GetCarId())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("query production_car error by id: %v\n", in.GetCarId()))
	}

	prod, err := l.svcCtx.Production.FindOne(l.ctx, car.ProdId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("query production error by id: %v\n", car.ProdId))
	}

	return &production.CarInfoResp{
		Id:        car.Id,
		ProdId:    prod.Id,
		Category:  prod.Category,
		Type:      car.Type,
		Series:    car.Series,
		Name:      car.Name,
		CreatedAt: car.CreatedAt,
		UpdatedAt: car.UpdatedAt,
	}, nil
}
