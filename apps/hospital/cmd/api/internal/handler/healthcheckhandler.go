package handler

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/logic"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHealthCheckLogic(r.Context(), svcCtx)
		resp, err := l.HealthCheck()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
