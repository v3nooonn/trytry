package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/product/internal/svc"
	"github.com/v3nooonn/trytry/apps/product/pb/product"

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

func (l *InfoLogic) Info(in *product.InfoReq) (*product.Production, error) {
	// todo: add your logic here and delete this line

	return &product.Production{}, nil
}
