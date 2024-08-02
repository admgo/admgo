package errorx

import (
	"context"
	"github.com/admgo/admgo/pkg/errorx/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
	"strconv"
)

func FromError(err error) *status.Status {
	err = errors.Cause(err)
	if s, ok := err.(RpcStatus); ok {
		grpcStatus, e := gRPCStatusFromXCode(s)
		if e == nil {
			return grpcStatus
		}
	}

	var grpcStatus *status.Status
	switch err {
	case context.Canceled:
		grpcStatus, _ = gRPCStatusFromXCode(RpcCanceled)
	case context.DeadlineExceeded:
		grpcStatus, _ = gRPCStatusFromXCode(RpcDeadline)
	default:
		grpcStatus, _ = status.FromError(err)
	}

	return grpcStatus
}
func gRPCStatusFromXCode(s RpcStatus) (*status.Status, error) {
	var sts *Status
	sts = Error(HttpErr{
		code: http.StatusOK,
		body: HttpErrBody{
			code: s.Code(),
			msg:  s.Message(),
		},
	})
	for _, detail := range s.Details() {
		if msg, ok := detail.(proto.Message); ok {
			_, _ = sts.WithDetails(msg)
		}
	}

	stas := status.New(codes.Unknown, strconv.Itoa(int(sts.Code())))
	return stas.WithDetails(sts.Proto())
}
func Error(herr HttpErr) *Status {
	return &Status{
		sts: &types.Status{
			Code:    herr.Body().code,
			Message: herr.Body().msg,
		},
	}
}

func GrpcStatusToXCode(gstatus *status.Status) HttpError {
	details := gstatus.Details()
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		if pb, ok := detail.(proto.Message); ok {
			return FromProto(pb)
		}
	}

	return toXCode(gstatus)
}
func FromProto(pbMsg proto.Message) HttpError {
	msg, ok := pbMsg.(*types.Status)
	if ok {
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return &HttpErr{
				code: http.StatusOK,
				body: HttpErrBody{
					msg: string(msg.Code),
				},
			}
		}
		return &HttpErr{
			code: http.StatusOK,
			body: HttpErrBody{
				code: msg.Code,
				msg:  msg.Message,
			}}
	}

	return ServerErr
}

func toXCode(grpcStatus *status.Status) *HttpErr {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.DeadlineExceeded:
		return HttpDeadline
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.Unknown:
		return String(grpcStatus.Message())
	}

	return ServerErr
}

//	func Errorf(code HttpErr, format string, args ...interface{}) *HttpErr {
//		code.Body().msg = fmt.Sprintf(format, args...)
//		return Error(code)
//	}
func (s *Status) WithDetails(msgs ...proto.Message) (*Status, error) {
	for _, msg := range msgs {
		anyMsg, err := anypb.New(msg)
		if err != nil {
			return s, err
		}
		s.sts.Details = append(s.sts.Details, anyMsg)
	}

	return s, nil
}
func String(s string) *HttpErr {
	if len(s) == 0 {
		return OK
	}
	code, err := strconv.Atoi(s)
	if err != nil {
		return ServerErr
	}

	return &HttpErr{code: int32(code)}
}
