package errorx

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func Handler(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		httpx.WriteJson(w, http.StatusOK, resp)
		return
	}

	var respErr any
	var c int

	sts, ok := status.FromError(err)
	if ok {
		c = int(HTTPStatusFromCode(sts.Code()))
		respErr = &ErrIns{
			CodeF:    int(HTTPStatusFromCode(sts.Code())),
			StatusF:  sts.Code().String(),
			MessageF: sts.Message(),
		}
	} else {
		if e, ok := err.(*ErrIns); ok {
			c = e.Code()
			respErr = &ErrIns{CodeF: e.Code(), StatusF: e.Status(), MessageF: e.Message()}
		} else {
			respErr = &ErrIns{CodeF: http.StatusInternalServerError, StatusF: http.StatusText(http.StatusInternalServerError), MessageF: err.Error()}
		}
	}

	logx.WithContext(r.Context()).Errorf("[ERROR] : %+v ", err)

	httpx.WriteJson(w, c, struct {
		Error any `json:"error"`
	}{
		Error: respErr,
	})
}
