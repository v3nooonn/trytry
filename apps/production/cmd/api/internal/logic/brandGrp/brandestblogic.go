package brandGrp

import (
	"context"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandEstbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandEstbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandEstbLogic {
	return &BrandEstbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandEstbLogic) BrandEstb(req *types.BrandEstbReq) (resp *types.BrandEstbResp, err error) {
	// todo: add your logic here and delete this line

	return
}
