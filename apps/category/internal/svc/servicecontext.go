package svc

import (
	"github.com/v3nooonn/trytry/apps/category/internal/config"
	"github.com/v3nooonn/trytry/apps/category/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	Category model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:   c,
		Category: model.NewCategoryModel(db, c.Cache),
	}
}
