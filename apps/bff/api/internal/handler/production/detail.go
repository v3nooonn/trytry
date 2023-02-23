package production

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/production"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	e "github.com/v3nooonn/trytry/pkg/errorx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := production.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()

		e.RespHandler(r, w, resp, err)
	}
}
