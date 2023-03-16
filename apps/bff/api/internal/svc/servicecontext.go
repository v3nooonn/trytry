package svc

import (
	"github.com/v3nooonn/trytry/apps/bff/api/internal/config"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/middleware"
	orgClient "github.com/v3nooonn/trytry/apps/organization/organizationclient"
	proClient "github.com/v3nooonn/trytry/apps/product/productclient"
	"github.com/v3nooonn/trytry/pkg/interceptor/outgoing"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	// configuration
	Config config.Config
	// middleware
	Public         rest.Middleware
	Authentication rest.Middleware
	Authorization  rest.Middleware
	Language       rest.Middleware
	RemoteAddr     rest.Middleware
	// rpc
	OrganizationRPC orgClient.Organization
	ProductionRPC   proClient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		// configuration
		Config: c,
		// middleware
		Public:         middleware.NewPublicMiddleware().Handle,
		Authentication: middleware.NewAuthenticationMiddleware().Handle,
		Authorization:  middleware.NewAuthorizationMiddleware().Handle,
		Language:       middleware.NewLanguageMiddleware().Handle,
		RemoteAddr:     middleware.NewRemoteAddrMiddleware().Handle,
		// rpc
		OrganizationRPC: orgClient.NewOrganization(zrpc.MustNewClient(
			c.Organization,
			zrpc.WithUnaryClientInterceptor(outgoing.FakeOut)),
		),
		ProductionRPC: proClient.NewProduct(zrpc.MustNewClient(
			c.Production,
			zrpc.WithUnaryClientInterceptor(outgoing.FakeOut)),
		),
	}
}
