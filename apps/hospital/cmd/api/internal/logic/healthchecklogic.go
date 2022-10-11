package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck() (resp *types.HealthResp, err error) {
	resp = &types.HealthResp{Up: "ok"}

	return
}
