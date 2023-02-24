package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductionCarModel = (*customProductionCarModel)(nil)

type (
	ProductionCarModel interface {
		productionCarModel
	}

	customProductionCarModel struct {
		*defaultProductionCarModel
	}
)

func NewProductionCarModel(conn sqlx.SqlConn, c cache.CacheConf) ProductionCarModel {
	return &customProductionCarModel{
		defaultProductionCarModel: newProductionCarModel(conn, c),
	}
}

//func (c customProductionCarModel) FindByProdID(ctx context.Context, prodID int64) ([]*ProductionCar, error) {
//	sql, args, err := c.Columns(ctx).Where(squirrel.Eq{"prod_id": prodID}).OrderBy("id DESC").ToSql()
//	if err != nil {
//		return nil, errors.Wrap(err, "build production_cars by prod_id sql error")
//	}
//
//	resp := make([]*ProductionCar, 0)
//	if err := c.QueryRowsNoCacheCtx(ctx, &resp, sql, args...); err != nil {
//		return nil, errors.Wrap(err, "query production_cars by prod_id error")
//	}
//
//	return resp, nil
//}
