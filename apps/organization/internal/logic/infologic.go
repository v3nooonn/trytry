package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/organization/internal/svc"
	"github.com/v3nooonn/trytry/apps/organization/pb/organization"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *organization.InfoReq) (*organization.InfoResp, error) {
	// todo: add your logic here and delete this line

	return &organization.InfoResp{}, nil
}
