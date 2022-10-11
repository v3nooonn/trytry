package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/v3nooonn/trytry/apps/oauth/cmd/rpc/oauthclient"
	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationMiddleware struct {
	Config   config.Config
	OAuthRpc oauthclient.Oauth
}

func NewAuthorizationMiddleware(c config.Config, rpc oauthclient.Oauth) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		Config:   c,
		OAuthRpc: rpc,
	}
}

func (m *AuthorizationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("--- Production:Middleware: Authorization Before")

		resp, err := m.OAuthRpc.Authorization(r.Context(), &oauthclient.Request{})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(MiddlewareError{
				Code:    http.StatusForbidden,
				Message: http.StatusText(http.StatusForbidden),
			})
			return
		}
		if !resp.Permitted {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(MiddlewareError{
				Code:    http.StatusForbidden,
				Message: http.StatusText(http.StatusForbidden),
			})
			return
		}

		next(w, r)

		logx.Info("--- Production:Middleware: Authorization After")
	}
}
