package session

import (
	"context"

	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSingleSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSingleSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSingleSessionLogic {
	return &GetSingleSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSingleSessionLogic) GetSingleSession(req *types.SessionReq) (resp *types.SessionRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
