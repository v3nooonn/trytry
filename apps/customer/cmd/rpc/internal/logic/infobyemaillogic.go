package logic

import (
	"context"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/customer/cmd/rpc/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/customer/cmd/rpc/pb/customer"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoByEmailLogic {
	return &InfoByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoByEmailLogic) InfoByEmail(in *customer.InfoByEmailReq) (*customer.InfoResp, error) {
	c, err := l.svcCtx.CustomerModel.FindByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}

	return &customer.InfoResp{
		Id:     uint64(c.Id),
		Name:   c.Name,
		Email:  c.Email,
		Secret: c.Secret,
	}, nil
}
