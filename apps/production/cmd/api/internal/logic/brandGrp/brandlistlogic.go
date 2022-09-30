package brandGrp

import (
	"context"
	"net/http"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandListLogic) BrandList(req *types.BrandListReq) (resp *types.BrandListResp, err error) {
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
