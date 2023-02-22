package dashboard

import (
	"context"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HumanServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHumanServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HumanServerLogic {
	return &HumanServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HumanServerLogic) HumanServer() (resp *types.HumServerResp, err error) {

	return
}
