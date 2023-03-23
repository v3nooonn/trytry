package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/brand/internal/svc"
	"github.com/v3nooonn/trytry/apps/brand/pb/brand"

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

func (l *InfoLogic) Info(in *brand.InfoReq) (*brand.BrandInfo, error) {
	b, err := l.svcCtx.Brand.FindOne(l.ctx, in.BrandId)
	if err != nil {
		return nil, err
	}

	return &brand.BrandInfo{
		Id:           b.Id,
		OwnBy:        b.OwnBy,
		CategoryId:   b.CategoryId,
		Name:         b.Name,
		Abbreviation: b.Abbreviation,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}, nil
}
