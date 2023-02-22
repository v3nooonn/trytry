package ping

import (
	"net/http"

	e "github.com/levvel-health/rpms-service/pkg/error"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/ping"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ping.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()

		e.RespHandler(r, w, resp, err)
	}
}
