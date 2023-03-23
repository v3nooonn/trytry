// Code generated by goctl. DO NOT EDIT.
// Source: category.proto

package server

import (
	"context"

	"github.com/v3nooonn/trytry/apps/category/internal/logic"
	"github.com/v3nooonn/trytry/apps/category/internal/svc"
	"github.com/v3nooonn/trytry/apps/category/pb/category"
)

type CategoryServer struct {
	svcCtx *svc.ServiceContext
	category.UnimplementedCategoryServer
}

func NewCategoryServer(svcCtx *svc.ServiceContext) *CategoryServer {
	return &CategoryServer{
		svcCtx: svcCtx,
	}
}

func (s *CategoryServer) Info(ctx context.Context, in *category.InfoReq) (*category.CategoryInfo, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}