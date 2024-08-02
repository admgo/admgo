package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserNameLogic {
	return &UpdateUserNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserName 更新用户名称
func (l *UpdateUserNameLogic) UpdateUserName(in *pb.UpdateUserNameRequest) (*pb.UpdateUserNameResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserNameResponse{}, nil
}
