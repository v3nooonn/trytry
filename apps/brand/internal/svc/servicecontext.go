package svc

import (
	"github.com/v3nooonn/trytry/apps/brand/internal/config"
	"github.com/v3nooonn/trytry/apps/brand/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Brand  model.BrandModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		Brand:  model.NewBrandModel(db, c.Cache),
	}
}
