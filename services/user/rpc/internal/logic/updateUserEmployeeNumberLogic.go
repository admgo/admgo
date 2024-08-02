package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserEmployeeNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserEmployeeNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserEmployeeNumberLogic {
	return &UpdateUserEmployeeNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserEmployeeNumber 更新用户员工号
func (l *UpdateUserEmployeeNumberLogic) UpdateUserEmployeeNumber(in *pb.UpdateUserEmployeeNumberRequest) (*pb.UpdateUserEmployeeNumberResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserEmployeeNumberResponse{}, nil
}
