package error

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RespHandler(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		httpx.WriteJson(w, http.StatusOK, resp)
		return
	}

	var respErr any
	var c int

	sts, ok := status.FromError(err)
	if ok {
		c = int(HTTPStatusFromCode(sts.Code()))
		respErr = &ErrStc{
			CodeF:    int(HTTPStatusFromCode(sts.Code())),
			StatusF:  sts.Code().String(),
			MessageF: sts.Message(),
		}
	} else {
		if e, ok := err.(*ErrStc); ok {
			c = e.Code()
			respErr = &ErrStc{CodeF: e.Code(), StatusF: e.Status(), MessageF: e.Message()}
		} else if e, ok := err.(*ErrStc); ok {
			c = e.Code()
			respErr = &ErrStc{CodeF: e.Code(), StatusF: e.Status(), MessageF: e.Message()}
		} else {
			respErr = &ErrStc{CodeF: http.StatusInternalServerError, StatusF: http.StatusText(http.StatusInternalServerError), MessageF: err.Error()}
		}
	}

	logx.WithContext(r.Context()).Errorf("[ERROR] : %+v ", err)

	httpx.WriteJson(w, c, struct {
		Error any `json:"error"`
	}{
		Error: respErr,
	})
}

type ErrStc struct {
	CodeF    int    `json:"code"`
	StatusF  string `json:"status"`
	MessageF string `json:"message"`
}

func (r *ErrStc) Code() int {
	return r.CodeF
}

func (r *ErrStc) Status() string {
	return r.StatusF
}

func (r *ErrStc) Message() string {
	return r.MessageF
}

func (r *ErrStc) Error() string {
	return fmt.Sprintf("code: %v, status: %s, message: %s", r.Code(), r.Status(), r.Message())
}

func newErr(c int, s string, m string) *ErrStc {
	return &ErrStc{
		CodeF:    c,
		StatusF:  s,
		MessageF: m,
	}
}

func BadRequest(m string) error {
	return newErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), m)
}

func Unauthorized(m string) *ErrStc {
	return newErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), m)
}

func Forbidden(m string) *ErrStc {
	return newErr(http.StatusForbidden, http.StatusText(http.StatusForbidden), m)
}

func NotFound(m string) *ErrStc {
	return newErr(http.StatusNotFound, http.StatusText(http.StatusNotFound), m)
}

func Internal(m string) *ErrStc {
	return newErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), m)
}

func AlreadyExists(m string) error {
	return newErr(http.StatusConflict, http.StatusText(http.StatusConflict), m)
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) uint32 {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists, codes.Aborted:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Canceled:
		return 499
	case codes.Unknown, codes.Internal, codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}
