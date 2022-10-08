package svc

import "github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
