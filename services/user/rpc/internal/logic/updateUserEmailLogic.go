package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserEmailLogic {
	return &UpdateUserEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserEmail 更新用户邮箱
func (l *UpdateUserEmailLogic) UpdateUserEmail(in *pb.UpdateUserEmailRequest) (*pb.UpdateUserEmailResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserEmailResponse{}, nil
}
