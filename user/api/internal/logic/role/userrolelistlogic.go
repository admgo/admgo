package role

import (
	"context"

	"github.com/admgo/admgo/user/api/internal/svc"
	"github.com/admgo/admgo/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleListLogic {
	return &UserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleListLogic) UserRoleList() (resp []types.UserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
