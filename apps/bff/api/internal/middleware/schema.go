package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/v3nooonn/trytry/pkg/constant/types"
	e "github.com/v3nooonn/trytry/pkg/errorx"
	"github.com/v3nooonn/trytry/pkg/utils/contextx"
)

type SchemaMiddleware struct{}

func NewSchemaMiddleware() *SchemaMiddleware {
	return &SchemaMiddleware{}
}

func (m *SchemaMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		encoded := r.Header.Get(types.CtxKeyTenant.Key())
		if len(encoded) == 0 {
			e.HandlerFuncErr(w, r, e.BadRequest("missing required headers"))
			return
		}

		bytes, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			e.HandlerFuncErr(w, r, e.Internal(err.Error()))
			return
		}
		tnt := new(contextx.Tenant)
		if err := json.Unmarshal(bytes, tnt); err != nil {
			e.HandlerFuncErr(w, r, e.Internal(err.Error()))
			return
		}

		next(w, r.WithContext(
			contextx.WithContext(r.Context()).
				WithValue(types.CtxKeyTenant, tnt).
				WithValue(types.CtxKeyTenantSchema, tnt.Schema).Context(),
		))
	}
}
