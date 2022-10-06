package svc

import (
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/config"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/model/brand"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/model/car"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	// configuration
	Config config.Config

	// models
	ProductionBrandModel brand.ProductionBrandModel
	ProductionCarModel   car.ProductionCarModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("mysql", c.DB.DataSource)
	return &ServiceContext{
		Config:               c,
		ProductionBrandModel: brand.NewProductionBrandModel(conn, c.Cache),
		ProductionCarModel:   car.NewProductionCarModel(conn, c.Cache),
	}
}
