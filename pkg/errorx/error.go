package error

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

type ErrIns struct {
	CodeF    int    `json:"code"`
	StatusF  string `json:"status"`
	MessageF string `json:"message"`
}

func (r *ErrIns) Code() int {
	return r.CodeF
}

func (r *ErrIns) Status() string {
	return r.StatusF
}

func (r *ErrIns) Message() string {
	return r.MessageF
}

func (r *ErrIns) Error() string {
	return fmt.Sprintf("code: %v, status: %s, message: %s", r.Code(), r.Status(), r.Message())
}

func newErr(c int, s string, m string) *ErrIns {
	return &ErrIns{
		CodeF:    c,
		StatusF:  s,
		MessageF: m,
	}
}

func BadRequest(m string) error {
	return newErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), m)
}

func Unauthorized(m string) *ErrIns {
	return newErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), m)
}

func Forbidden(m string) *ErrIns {
	return newErr(http.StatusForbidden, http.StatusText(http.StatusForbidden), m)
}

func NotFound(m string) *ErrIns {
	return newErr(http.StatusNotFound, http.StatusText(http.StatusNotFound), m)
}

func AlreadyExists(m string) error {
	return newErr(http.StatusConflict, http.StatusText(http.StatusConflict), m)
}

func Internal(m string) *ErrIns {
	return newErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), m)
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
