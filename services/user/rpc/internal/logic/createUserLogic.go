package logic

import (
	"context"
	"errors"
	"github.com/admgo/admgo/services/user/rpc/internal/code"
	"github.com/admgo/admgo/services/user/rpc/internal/db/models"
	"gorm.io/gorm"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 创建用户
func (l *CreateUserLogic) CreateUser(in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	err := l.transaction(func(tx *gorm.DB) error {
		err := l.svcCtx.UserModel.WithTransaction(tx).Insert(l.ctx, &models.User{
			UserName: in.UserName,
			Name:     in.Name,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return code.MysqlErr
			}
			return err
		}
		err = l.svcCtx.TokenModel.WithTransaction(tx).Insert(l.ctx, &models.Token{
			UserName: in.UserName,
			Name:     in.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return &pb.CreateUserResponse{}, err
}

func (l *CreateUserLogic) transaction(fn func(tx *gorm.DB) error) error {
	return l.svcCtx.DB.Transaction(fn)
}
