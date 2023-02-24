package error

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HandlerFuncErr(w http.ResponseWriter, r *http.Request, err error) {
	e2, ok := err.(*ErrIns)
	if !ok {
		e2 = Internal(err.Error())
	}

	logx.WithContext(r.Context()).Errorf("[ERROR] : %+v ", err)

	//w.WriteHeader(e2.Code())
	//w.Header().Add(types.HeaderFieldContentType.String(), content.ApplicationJSON)
	//bytes, _ := json.Marshal(e2)
	//_, _ = w.Write(bytes)
	httpx.WriteJson(w, e2.Code(), struct {
		Error any `json:"error"`
	}{
		Error: e2,
	})

}
