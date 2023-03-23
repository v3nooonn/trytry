package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		Pagination(ctx context.Context, cursor, limit int64) ([]Product, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c),
	}
}

func (c customProductModel) Pagination(ctx context.Context, cursor, limit int64) ([]Product, error) {
	sql, _, err := squirrel.Select(productRows).
		From(c.tableName()).
		Where(squirrel.Gt{"id": cursor}).
		Limit(uint64(limit)).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "sql construction error")
	}

	prods := make([]Product, 0, limit)

	if err := c.QueryRowsNoCache(&prods, sql, cursor); err != nil {
		return nil, errors.Wrap(err, "production pagination error")
	}

	return prods, nil
}
