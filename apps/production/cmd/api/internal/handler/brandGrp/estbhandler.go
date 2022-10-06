package brandGrp

import (
	"net/http"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/logic/brandGrp"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/production/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EstbHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandEstbReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := brandGrp.NewEstbLogic(r.Context(), svcCtx)
		resp, err := l.Estb(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
