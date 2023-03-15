package dashboard

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/dashboard"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CustomerServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dashboard.NewCustomerServiceLogic(r.Context(), svcCtx)
		resp, err := l.CustomerService()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
