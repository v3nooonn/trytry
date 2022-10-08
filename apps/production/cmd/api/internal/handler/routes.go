// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	brandGrp "github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/handler/brandGrp"
	carGrp "github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/handler/carGrp"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authentication, serverCtx.Authorization},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/estb",
					Handler: carGrp.EstbHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: carGrp.InfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/production/car"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authentication, serverCtx.Authorization},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/estb",
					Handler: brandGrp.EstbHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: brandGrp.ListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/production/brand"),
	)
}
