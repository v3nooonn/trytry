// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	oauthGrp "github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/api/internal/handler/oauthGrp"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/access",
				Handler: oauthGrp.AccessHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/oauth"),
	)
}