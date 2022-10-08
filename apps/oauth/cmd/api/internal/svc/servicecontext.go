package svc

import (
	"github.com/v3nooonn/trytry-based-on-looklook/apps/customer/cmd/rpc/customerclient"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	// configuration
	Config config.Config

	// rpc
	CustomerRpc customerclient.Customer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CustomerRpc: customerclient.NewCustomer(zrpc.MustNewClient(c.CustomerRpc)),
	}
}
