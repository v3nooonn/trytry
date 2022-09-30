package car

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductionCarModel = (*customProductionCarModel)(nil)

type (
	// ProductionCarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductionCarModel.
	ProductionCarModel interface {
		productionCarModel
	}

	customProductionCarModel struct {
		*defaultProductionCarModel
	}
)

// NewProductionCarModel returns a model for the database table.
func NewProductionCarModel(conn sqlx.SqlConn) ProductionCarModel {
	return &customProductionCarModel{
		defaultProductionCarModel: newProductionCarModel(conn),
	}
}
