package types

type ContextKey string

func (r ContextKey) Key() string {
	return string(r)
}

const (
	CtxKeyMeta         ContextKey = "META_DATA_KEY"
	CtxKeyTenant       ContextKey = "TENANT"
	CtxKeyTenantSchema ContextKey = "TENANT_SCHEMA"
)
