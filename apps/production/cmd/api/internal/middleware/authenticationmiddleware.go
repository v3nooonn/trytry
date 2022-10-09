package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/config"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationMiddleware struct {
	Config config.Config
}

func NewAuthenticationMiddleware(c config.Config) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		Config: c,
	}
}

type MiddlewareError struct {
	Code    uint64
	Message string
}

func (m *AuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("--- Production:Middleware: Authentication Before")

		headerStr := r.Header.Get("Authorization")
		headers := strings.Split(headerStr, " ")
		if len(headers) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(MiddlewareError{
				Code:    http.StatusUnauthorized,
				Message: fmt.Sprintf("Invalid Authorization."),
			})
			return
		}
		tokenStr := headers[1]

		// std := jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(m.Config.OAuth.PublicKey))
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(MiddlewareError{
				Code:    http.StatusUnauthorized,
				Message: fmt.Sprintf("Invalid Authorization."),
			})
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(MiddlewareError{
				Code:    http.StatusUnauthorized,
				Message: fmt.Sprintf("Invalid Authorization."),
			})
			return
		}

		next(w, r)

		logx.Info("--- Production:Middleware: Authentication After")
	}
}
