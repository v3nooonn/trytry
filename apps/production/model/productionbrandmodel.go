package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductionBrandModel = (*customProductionBrandModel)(nil)

type (
	// ProductionBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductionBrandModel.
	ProductionBrandModel interface {
		productionBrandModel
	}

	customProductionBrandModel struct {
		*defaultProductionBrandModel
	}
)

// NewProductionBrandModel returns a model for the database table.
func NewProductionBrandModel(conn sqlx.SqlConn, c cache.CacheConf) ProductionBrandModel {
	return &customProductionBrandModel{
		defaultProductionBrandModel: newProductionBrandModel(conn, c),
	}
}
