package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BrandModel = (*customBrandModel)(nil)

type (
	// BrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBrandModel.
	BrandModel interface {
		brandModel
	}

	customBrandModel struct {
		*defaultBrandModel
	}
)

// NewBrandModel returns a model for the database table.
func NewBrandModel(conn sqlx.SqlConn, c cache.CacheConf) BrandModel {
	return &customBrandModel{
		defaultBrandModel: newBrandModel(conn, c),
	}
}
