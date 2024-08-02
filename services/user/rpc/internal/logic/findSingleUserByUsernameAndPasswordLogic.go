package logic

import (
	"context"
	"github.com/admgo/admgo/services/user/rpc/internal/code"
	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSingleUserByUsernameAndPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSingleUserByUsernameAndPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSingleUserByUsernameAndPasswordLogic {
	return &FindSingleUserByUsernameAndPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindSingleUserByUsernameAndPassword 根据用户名密码查找指定用户
func (l *FindSingleUserByUsernameAndPasswordLogic) FindSingleUserByUsernameAndPassword(in *pb.FindSingleUserByUsernameAndPasswordRequest) (*pb.FindSingleUserByUsernameAndPasswordResponse, error) {
	data, err := l.svcCtx.UserModel.FindByUserIDAndPassword(l.ctx, in.UserName, in.Password)
	if err != nil {
		return nil, err
	}

	if data == nil {

		return &pb.FindSingleUserByUsernameAndPasswordResponse{}, code.IncorrectUserNameORPassword
	}

	return &pb.FindSingleUserByUsernameAndPasswordResponse{
		UserID:         data.ID,
		Name:           data.Name,
		UserName:       data.UserName,
		Phone:          data.Phone,
		Email:          data.Email,
		EmployeeNumber: data.EmployeeNumber,
		Avatar:         data.Avatar,
		IsAdmin:        data.IsAdmin,
	}, nil
}
