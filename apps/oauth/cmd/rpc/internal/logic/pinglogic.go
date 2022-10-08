package logic

import (
	"context"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/rpc/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/rpc/pb/oauth"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *oauth.Request) (*oauth.Response, error) {
	// todo: add your logic here and delete this line

	return &oauth.Response{}, nil
}
