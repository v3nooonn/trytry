package middleware

import "net/http"

type RemoteAddrMiddleware struct{}

func NewRemoteAddrMiddleware() *RemoteAddrMiddleware {
	return &RemoteAddrMiddleware{}
}

func (m *RemoteAddrMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
