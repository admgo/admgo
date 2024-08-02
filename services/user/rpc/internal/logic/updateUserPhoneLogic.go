package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPhoneLogic {
	return &UpdateUserPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserPhone 更新用户手机号
func (l *UpdateUserPhoneLogic) UpdateUserPhone(in *pb.UpdateUserPhoneRequest) (*pb.UpdateUserPhoneResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserPhoneResponse{}, nil
}
