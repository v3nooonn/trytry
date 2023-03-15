package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrganizationModel = (*customOrganizationModel)(nil)

type (
	// OrganizationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrganizationModel.
	OrganizationModel interface {
		organizationModel
	}

	customOrganizationModel struct {
		*defaultOrganizationModel
	}
)

// NewOrganizationModel returns a model for the database table.
func NewOrganizationModel(conn sqlx.SqlConn, c cache.CacheConf) OrganizationModel {
	return &customOrganizationModel{
		defaultOrganizationModel: newOrganizationModel(conn, c),
	}
}
