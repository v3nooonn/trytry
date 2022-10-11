package carGrp

import (
	"context"
	"net/http"

	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.CarInfoReq) (resp *types.CarInfoResp, err error) {
	pc, err := l.svcCtx.ProductionCarModel.FindOne(l.ctx, int64(req.ID))
	switch err {
	case nil:
		return &types.CarInfoResp{
			Code:    http.StatusOK,
			Message: "ok",
			Data: types.Car{
				ID:    uint64(pc.Id),
				Brand: pc.Brand,
				Serie: pc.Serie,
			},
		}, nil
	case sqlc.ErrNotFound:
		return &types.CarInfoResp{
			Code:    http.StatusOK,
			Message: "Not Found",
			Data:    types.Car{},
		}, nil
	default:
		return nil, err
	}
}
