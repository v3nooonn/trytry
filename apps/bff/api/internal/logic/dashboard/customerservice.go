package dashboard

import (
	"context"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerServiceLogic {
	return &CustomerServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomerServiceLogic) CustomerService() (resp *types.CustomerServiceResp, err error) {

	return
}
