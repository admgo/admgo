package errorx

import (
	"net/http"
	"strconv"
)

type (
	HttpError interface {
		Errorx
		Code() int32        //http状态码
		Body() *HttpErrBody //错误主体
	}
	HttpErrBody struct {
		code int32
		msg  string
	}
	HttpErr struct {
		code int32
		body HttpErrBody
	}
)

func (e HttpErr) Error() string {
	if len(e.body.msg) > 0 {
		return e.body.msg
	}

	return strconv.Itoa(int(e.code))
}

func (e HttpErr) Code() int32 {
	return e.code
}
func (e HttpErr) Body() *HttpErrBody {
	return &HttpErrBody{
		code: e.body.code,
		msg:  e.body.msg,
	}
}

func (e HttpErr) ErrTag() string {
	return "httperror"
}

func NewHttpErr(bizcode int32, msg string) *HttpErr {
	return &HttpErr{
		code: http.StatusOK,
		body: HttpErrBody{
			code: bizcode,
			msg:  msg,
		},
	}
}

func NewHttpErrAndSetHttpCode(httpcode, bizcode int32, msg string) *HttpErr {
	return &HttpErr{
		code: httpcode,
		body: HttpErrBody{
			code: bizcode,
			msg:  msg,
		},
	}
}
