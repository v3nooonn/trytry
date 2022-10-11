package svc

import (
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/config"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	Authentication rest.Middleware
	Authorization  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Authentication: middleware.NewAuthenticationMiddleware().Handle,
		Authorization:  middleware.NewAuthorizationMiddleware().Handle,
	}
}
