package ping

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/ping"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	e "github.com/v3nooonn/trytry/pkg/errorx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ping.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()

		e.HandlerErr(r, w, resp, err)
	}
}
