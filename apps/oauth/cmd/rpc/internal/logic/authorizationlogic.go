package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/oauth/cmd/rpc/internal/svc"
	"github.com/v3nooonn/trytry/apps/oauth/cmd/rpc/pb/oauth"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizationLogic {
	return &AuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthorizationLogic) Authorization(in *oauth.Request) (*oauth.Response, error) {
	logx.Info("--- OAuth:RPC: Authorization")

	return &oauth.Response{
		Permitted: true,
	}, nil
}
