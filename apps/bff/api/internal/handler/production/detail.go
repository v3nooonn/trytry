package production

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/production"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	e "github.com/v3nooonn/trytry/pkg/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailReq
		if err := httpx.Parse(r, &req); err != nil {
			e.HandlerErr(r, w, nil, e.BadRequest(err.Error()))
			return
		}

		l := production.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)

		e.HandlerErr(r, w, resp, err)
	}
}
