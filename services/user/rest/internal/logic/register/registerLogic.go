package register

import (
	"context"
	"github.com/admgo/admgo/services/user/rpc/user"

	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRsp, err error) {
	_, err = l.svcCtx.UserRPC.CreateUser(l.ctx, &user.CreateUserRequest{
		Name:     "34",
		UserName: "12",
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterRsp{}, nil
}
