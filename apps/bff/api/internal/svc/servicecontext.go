package svc

import (
	"github.com/v3nooonn/trytry/apps/bff/api/internal/config"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/middleware"
	"github.com/v3nooonn/trytry/pkg/interceptor/outgoing"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	pbProdSvc "github.com/v3nooonn/trytry/apps/production/productionservice"
)

type ServiceContext struct {
	Config config.Config

	// middleware
	Schema         rest.Middleware
	Authentication rest.Middleware
	Authorization  rest.Middleware
	Language       rest.Middleware
	RemoteAddr     rest.Middleware
	Public         rest.Middleware

	// RPC
	ProductionService pbProdSvc.ProductionService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Schema:         middleware.NewSchemaMiddleware().Handle,
		Authentication: middleware.NewAuthenticationMiddleware().Handle,
		Authorization:  middleware.NewAuthorizationMiddleware().Handle,
		Language:       middleware.NewLanguageMiddleware().Handle,
		RemoteAddr:     middleware.NewRemoteAddrMiddleware().Handle,
		Public:         middleware.NewPublicMiddleware().Handle,

		ProductionService: pbProdSvc.NewProductionService(zrpc.MustNewClient(
			c.ProdService,
			zrpc.WithUnaryClientInterceptor(outgoing.Tenant),
		)),
	}
}
