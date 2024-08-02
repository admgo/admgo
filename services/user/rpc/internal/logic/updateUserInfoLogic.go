package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserInfo 更新用户信息（包括邮箱、用户名称等）
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserInfoResponse{}, nil
}
