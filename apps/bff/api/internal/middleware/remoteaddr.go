package middleware

import "net/http"

type RemoteAddrMiddleware struct {
}

func NewRemoteAddrMiddleware() *RemoteAddrMiddleware {
	return &RemoteAddrMiddleware{}
}

func (m *RemoteAddrMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
