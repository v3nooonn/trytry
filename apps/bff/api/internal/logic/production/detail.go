package production

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	pbSvc "github.com/v3nooonn/trytry/apps/production/productionservice"
	e "github.com/v3nooonn/trytry/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailReq) (*types.Detail, error) {
	fmt.Println("cao")
	detail, err := l.svcCtx.ProductionService.CarInfo(l.ctx, &pbSvc.CarInfoReq{CarId: req.ID})
	if err != nil {
		return nil, err
	}

	resp := new(types.Detail)
	if err := copier.Copy(resp, detail); err != nil {
		return nil, errors.Wrap(e.Internal(err.Error()), "")
	}

	return resp, nil
}
