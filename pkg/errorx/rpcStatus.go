package errorx

import (
	"github.com/admgo/admgo/pkg/errorx/types"
	"github.com/golang/protobuf/ptypes"
	"strconv"
)

type (
	RpcStatus interface {
		Errorx
		Code() int32
		Message() string
		Details() []interface{}
	}
	Status struct {
		sts *types.Status
	}
)

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Code() int32 {
	return s.sts.Code
}

func (s *Status) Message() string {
	if s.sts.Message == "" {
		return strconv.Itoa(int(s.sts.Code))
	}

	return s.sts.Message
}

func (s *Status) Details() []interface{} {
	if s == nil || s.sts == nil {
		return nil
	}
	details := make([]interface{}, 0, len(s.sts.Details))
	for _, d := range s.sts.Details {
		detail := &ptypes.DynamicAny{}

		if err := d.UnmarshalTo(detail); err != nil {
			details = append(details, err)
			continue
		}
		details = append(details, detail.Message)
	}

	return details
}

func (s *Status) Proto() *types.Status {
	return s.sts
}
func (s *Status) ErrTag() string {
	return "rpcstatus"
}

func NewRpcStatus(code int32, msg string) *Status {
	return &Status{
		sts: &types.Status{
			Code:    code,
			Message: msg,
		},
	}
}
