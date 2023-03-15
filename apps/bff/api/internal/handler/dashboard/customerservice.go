package dashboard

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/dashboard"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	e "github.com/v3nooonn/trytry/pkg/expands/errorx"
)

func CustomerServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dashboard.NewCustomerServiceLogic(r.Context(), svcCtx)
		resp, err := l.CustomerService()
		e.Handler(r, w, resp, err)
	}
}
