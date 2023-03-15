package product

import (
	"net/http"

	"github.com/v3nooonn/trytry/apps/bff/api/internal/logic/product"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/svc"
	"github.com/v3nooonn/trytry/apps/bff/api/internal/types"
	e "github.com/v3nooonn/trytry/pkg/expands/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			e.Handler(r, w, nil, err)
			return
		}

		l := product.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		e.Handler(r, w, resp, err)
	}
}
