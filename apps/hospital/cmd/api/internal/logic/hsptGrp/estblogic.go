package hsptGrp

import (
	"context"

	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/types"

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

func (l *EstbLogic) Estb(req *types.HsptEstbReq) (resp *types.HsptEstbResp, err error) {
	resp = &types.HsptEstbResp{
		Name: "HHHHHHHHH",
	}

	return
}
