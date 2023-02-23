package contextx

import (
	"context"
)

type H struct {
	context context.Context
}

func WithContext(ctx context.Context) *H {
	return &H{context: ctx}
}

func (h *H) WithValue(k types.ContextKey, v any) *H {
	h.setContext(context.WithValue(h.Context(), k, v))
	return h
}

func (h *H) Context() context.Context {
	return h.context
}

func (h *H) setContext(ctx context.Context) *H {
	h.context = ctx
	return h
}
