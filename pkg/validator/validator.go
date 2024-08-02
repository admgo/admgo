package validator

import (
	"github.com/admgo/admgo/pkg/errorx"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type V struct {
	validator *validator.Validate
}

func NewValidate() *V {
	return &V{
		validator: validator.New(),
	}
}

func (v *V) Validate(_ *http.Request, data any) error {

	err := v.validator.Struct(data)
	if err != nil {
		return errorx.NewHttpErr(40001, err.Error())
	}
	return nil
}
