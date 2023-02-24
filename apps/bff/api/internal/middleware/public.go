package middleware

import "net/http"

type PublicMiddleware struct{}

func NewPublicMiddleware() *PublicMiddleware {
	return &PublicMiddleware{}
}

func (m *PublicMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
