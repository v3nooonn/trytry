package middleware

import "net/http"

type AllowedMiddleware struct{}

func NewAllowedMiddleware() *AllowedMiddleware {
	return &AllowedMiddleware{}
}

func (m *AllowedMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if it's required
		next(w, r)
	}
}
