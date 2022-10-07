package svc

import (
	"github.com/v3nooonn/trytry-based-on-looklook/apps/customer/cmd/rpc/internal/config"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/customer/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	CustomerModel model.CustomerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("mysql", c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		CustomerModel: model.NewCustomerModel(conn, c.Cache),
	}
}
