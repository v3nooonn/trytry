package ping

import (
	"context"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Pong, err error) {
	// todo: add your logic here and delete this line

	return
}
