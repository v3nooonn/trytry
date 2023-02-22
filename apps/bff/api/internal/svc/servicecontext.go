package svc

import (
	"github.com/v3nooonn/trytry/apps/bff/api/internal/config"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	lang           rest.Middleware
	remoteAddr     rest.Middleware
	allowed        rest.Middleware
	authentication rest.Middleware
	authorization  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		lang:           middleware.NewLangMiddleware().Handle,
		remoteAddr:     middleware.NewRemoteAddrMiddleware().Handle,
		allowed:        middleware.NewAllowedMiddleware().Handle,
		authentication: middleware.NewAuthenticationMiddleware().Handle,
		authorization:  middleware.NewAuthorizationMiddleware().Handle,
	}
}
