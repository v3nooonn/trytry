package carGrp

import (
	"context"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarEstbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarEstbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarEstbLogic {
	return &CarEstbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CarEstbLogic) CarEstb(req *types.CarEstbReq) (resp *types.CarEstbResp, err error) {
	// todo: add your logic here and delete this line

	return
}
