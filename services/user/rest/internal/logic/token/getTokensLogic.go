package token

import (
	"context"

	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokensLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokensLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokensLogic {
	return &GetTokensLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokensLogic) GetTokens(req *types.TokenReq) (resp *types.TokenRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
