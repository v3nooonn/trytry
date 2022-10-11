package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/customer/cmd/rpc/internal/svc"
	"github.com/v3nooonn/trytry/apps/customer/cmd/rpc/pb/customer"

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

func (l *InfoLogic) Info(in *customer.InfoReq) (*customer.InfoResp, error) {
	c, err := l.svcCtx.CustomerModel.FindOne(l.ctx, 1)
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
