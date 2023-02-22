package middleware

import "net/http"

type LangMiddleware struct{}

func NewLangMiddleware() *LangMiddleware {
	return &LangMiddleware{}
}

func (m *LangMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
