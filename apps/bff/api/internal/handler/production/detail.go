package production

import (
	"net/http"

	e "github.com/levvel-health/rpms-service/pkg/error"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/production"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := production.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()

		e.RespHandler(r, w, resp, err)
	}
}
