package svc

import (
	"github.com/v3nooonn/trytry/apps/product/internal/config"
	"github.com/v3nooonn/trytry/apps/product/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	Category model.CategoryModel
	Brand    model.BrandModel
	Product  model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//db := sqlx.NewSqlConn("postgres", c.DB.DataSource)
	db := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:   c,
		Category: model.NewCategoryModel(db, c.Cache),
		Brand:    model.NewBrandModel(db, c.Cache),
		Product:  model.NewProductModel(db, c.Cache),
	}
}
