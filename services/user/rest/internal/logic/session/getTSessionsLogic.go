package session

import (
	"context"

	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTSessionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTSessionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTSessionsLogic {
	return &GetTSessionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTSessionsLogic) GetTSessions(req *types.SessionReq) (resp *types.SessionRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
