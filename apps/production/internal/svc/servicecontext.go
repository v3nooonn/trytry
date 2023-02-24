package svc

import (
	"github.com/v3nooonn/trytry/apps/production/internal/config"
	"github.com/v3nooonn/trytry/apps/production/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	Production    model.ProductionModel
	ProductionCar model.ProductionCarModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("postgres", c.DB.DataSource)

	return &ServiceContext{
		Config:        c,
		Production:    model.NewProductionModel(db, c.Cache),
		ProductionCar: model.NewProductionCarModel(db, c.Cache),
	}
}
