package carGrp

import (
	"context"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarInfoLogic {
	return &CarInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CarInfoLogic) CarInfo(req *types.CarInfoReq) (resp *types.CarInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
