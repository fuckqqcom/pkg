package codex

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"sync/atomic"
)

var (
	_messages atomic.Value           // NOTE: stored map[int]string
	_codes    = map[int64]struct{}{} // register codes.
)

func Register(cm map[int64]string) {
	_messages.Store(cm)
}

type Code int64

func New(e int64) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e int64) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Code(e)
}

type Codes interface {
	Error() string
	Code() int64
	Add(int64) Code
	Message() string
	Details() []interface{}
	WithMsg(s string) Code
}

func (e Code) WithMsg(s string) Code {
	return errors.New(s).(Code)
}
func (e Code) Add(i int64) Code {
	return Code(i)
}

func (e Code) Error() string {
	return strconv.Itoa(int(e))
}

func (e Code) Code() int64 { return int64(e) }

func (e Code) Message() string {
	if cm, ok := _messages.Load().(map[int64]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return e.Error()
}

func ErrToCode(err error) Code {
	if err == nil {
		return OK
	}
	s, ok := status.FromError(err)
	if ok {
		return Code(s.Code())
	}
	return ServerErrCode
}

func StatusFromGrpcStatus(err codes.Code) int {
	switch err {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusGatewayTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
