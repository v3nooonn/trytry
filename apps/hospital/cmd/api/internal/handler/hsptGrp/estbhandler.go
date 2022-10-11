package hsptGrp

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/logic/hsptGrp"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EstbHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HsptEstbReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := hsptGrp.NewEstbLogic(r.Context(), svcCtx)
		resp, err := l.Estb(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
