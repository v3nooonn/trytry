package contextx

import (
	"context"

	"github.com/v3nooonn/trytry/pkg/constant/types"
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

// Tenant object of tenant in context
type Tenant struct {
	ID     int64
	Schema string
	Name   string
	Domain string
}

//func GetTen(ctx context.Context) Tenant {
//	tnt, _ := ctx.Value(types.CtxKeyTenantSchema).(*Tenant)
//	return *tnt
//}
//
//func SetTen(ctx context.Context, tenant Tenant) context.Context {
//	return context.WithValue(ctx, types.CtxKeyTenantSchema, tenant)
//}

func GetSchema(ctx context.Context) string {
	schema, _ := ctx.Value(types.CtxKeyTenantSchema).(string)
	return schema
}

func SetSchema(ctx context.Context, schema string) context.Context {
	return context.WithValue(ctx, types.CtxKeyTenantSchema, schema)
}
