package token

import (
	"context"

	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTokenLogic {
	return &AddTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTokenLogic) AddToken(req *types.TokenReq) (resp *types.TokenRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
