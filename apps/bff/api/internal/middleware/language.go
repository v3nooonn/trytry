package middleware

import "net/http"

type LanguageMiddleware struct{}

func NewLanguageMiddleware() *LanguageMiddleware {
	return &LanguageMiddleware{}
}

func (m *LanguageMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
