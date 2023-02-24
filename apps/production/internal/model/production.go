package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductionModel = (*customProductionModel)(nil)

type (
	ProductionModel interface {
		productionModel
		List(ctx context.Context) ([]*Production, error)
	}

	customProductionModel struct {
		*defaultProductionModel
	}
)

func NewProductionModel(conn sqlx.SqlConn, c cache.CacheConf) ProductionModel {
	return &customProductionModel{
		defaultProductionModel: newProductionModel(conn, c),
	}
}

func (c customProductionModel) List(ctx context.Context) ([]*Production, error) {
	sql, args, err := c.Columns(ctx).Where(squirrel.Eq{"status": 1}).OrderBy("id DESC").ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "production list error")
	}

	resp := make([]*Production, 0)
	if err := c.QueryRowsNoCacheCtx(ctx, &resp, sql, args...); err != nil {
		return nil, errors.Wrap(err, "production list error")
	}

	return resp, nil
}

