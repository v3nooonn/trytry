package svc

import (
	"github.com/v3nooonn/trytry/apps/product/internal/config"
	"github.com/v3nooonn/trytry/apps/product/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	Product model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:  c,
		Product: model.NewProductModel(db, c.Cache),
	}
}
