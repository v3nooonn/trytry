package logic

import (
	"context"

	"github.com/v3nooonn/trytry/apps/category/internal/svc"
	"github.com/v3nooonn/trytry/apps/category/pb/category"

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

func (l *InfoLogic) Info(in *category.InfoReq) (*category.CategoryInfo, error) {
	c, err := l.svcCtx.Category.FindOne(l.ctx, in.CategoryId)
	if err != nil {
		return nil, err
	}

	return &category.CategoryInfo{
		Id:           c.Id,
		Name:         c.Name,
		Abbreviation: c.Abbreviation,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}, nil
}
