package oauthGrp

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/oauth/cmd/api/internal/logic/oauthGrp"
	"github.com/v3nooonn/trytry/apps/oauth/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/oauth/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AccessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauthGrp.NewAccessLogic(r.Context(), svcCtx)
		resp, err := l.Access(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
