package brand

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductionBrandModel = (*customProductionBrandModel)(nil)

type (
	// ProductionBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductionBrandModel.
	ProductionBrandModel interface {
		productionBrandModel
		Pagination(ctx context.Context, page, size int64) ([]ProductionBrand, error)
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

// Pagination
func (m *defaultProductionBrandModel) Pagination(ctx context.Context, page, size int64) ([]ProductionBrand, error) {
	results := make([]ProductionBrand, 0)
	offset := int(size * (page - 1))

	query := fmt.Sprintf("select %s from %s limit %v offset %v", productionBrandRows, m.table, size, offset)

	if err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &results, query); err != nil {
		return nil, err
	}

	return results, nil
}
