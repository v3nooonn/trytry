package brandGrp

import (
	"context"
	"net/http"

	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.BrandListReq) (resp *types.BrandListResp, err error) {
	result, err := l.svcCtx.ProductionBrandModel.Pagination(l.ctx, int64(req.Page), int64(req.Size))
	if err != nil {
		return nil, err
	}

	resp = &types.BrandListResp{
		Code:    http.StatusOK,
		Message: "ok",
	}
	for _, pb := range result {
		resp.Data = append(resp.Data, types.Brand{
			ID:       uint64(pb.Id),
			Category: uint8(pb.Category),
			Name:     pb.Name,
		})
	}
	return
}
