package errorx

import (
	"github.com/admgo/admgo/pkg/errorx/types"
	"github.com/pkg/errors"
	"net/http"
)

func ErrHandler(err error) (int, any) {
	err = errors.Cause(err)
	if e, ok := err.(HttpErr); ok {
		return int(e.Code()), types.Status{
			Code:    e.Body().code,
			Message: e.Body().msg,
		}
	}
	if e, ok := err.(*HttpErr); ok {
		return int(e.Code()), types.Status{
			Code:    e.Body().code,
			Message: e.Body().msg,
		}
	}
	return http.StatusInternalServerError, err
}
