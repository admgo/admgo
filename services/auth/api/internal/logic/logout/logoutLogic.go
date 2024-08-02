package logout

import (
	"context"

	"github.com/admgo/admgo/services/auth/api/internal/svc"
	"github.com/admgo/admgo/services/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.LogoutRsp, err error) {
	// todo: add your logic here and delete this line

	return &types.LogoutRsp{}, nil
}
