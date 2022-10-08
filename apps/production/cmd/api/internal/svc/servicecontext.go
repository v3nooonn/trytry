package svc

import (
	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/rpc/oauthclient"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/config"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/middleware"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/model/brand"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/model/car"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	// configuration
	Config config.Config

	// rpc
	OAuthRpc oauthclient.Oauth

	// middlewares
	Authentication rest.Middleware
	Authorization  rest.Middleware

	// models
	ProductionBrandModel brand.ProductionBrandModel
	ProductionCarModel   car.ProductionCarModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rpcClient := oauthclient.NewOauth(zrpc.MustNewClient(c.OAuthRpc))

	conn := sqlx.NewSqlConn("mysql", c.DB.DataSource)
	return &ServiceContext{
		Config:               c,
		OAuthRpc:             rpcClient,
		Authentication:       middleware.NewAuthenticationMiddleware(c).Handle,
		Authorization:        middleware.NewAuthorizationMiddleware(c, rpcClient).Handle,
		ProductionBrandModel: brand.NewProductionBrandModel(conn, c.Cache),
		ProductionCarModel:   car.NewProductionCarModel(conn, c.Cache),
	}
}
