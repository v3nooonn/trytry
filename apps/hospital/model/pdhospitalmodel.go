package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PdHospitalModel = (*customPdHospitalModel)(nil)

type (
	// PdHospitalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPdHospitalModel.
	PdHospitalModel interface {
		pdHospitalModel
	}

	customPdHospitalModel struct {
		*defaultPdHospitalModel
	}
)

// NewPdHospitalModel returns a model for the database table.
func NewPdHospitalModel(conn sqlx.SqlConn, c cache.CacheConf) PdHospitalModel {
	return &customPdHospitalModel{
		defaultPdHospitalModel: newPdHospitalModel(conn, c),
	}
}
