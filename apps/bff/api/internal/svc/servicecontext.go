package svc

import (
	"github.com/v3nooonn/trytry/apps/bff/api/internal/config"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	Authentication rest.Middleware
	Authorization  rest.Middleware
	Language       rest.Middleware
	RemoteAddr     rest.Middleware
	Public         rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Authentication: middleware.NewAuthenticationMiddleware().Handle,
		Authorization:  middleware.NewAuthorizationMiddleware().Handle,
		Language:       middleware.NewLanguageMiddleware().Handle,
		RemoteAddr:     middleware.NewRemoteAddrMiddleware().Handle,
		Public:         middleware.NewPublicMiddleware().Handle,
	}
}
