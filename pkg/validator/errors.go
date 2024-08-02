package validator

import "github.com/admgo/admgo/pkg/errorx"

var (
	BadRequests = errorx.NewHttpErr(40001, "bad requests")
)

func NewXcode(msg string) {
	errorx.NewHttpErr(40001, "bad requests")
}
