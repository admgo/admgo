package logic

import (
	"context"

	"github.com/admgo/admgo/services/storager/api/internal/svc"
	"github.com/admgo/admgo/services/storager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoragerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoragerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoragerLogic {
	return &StoragerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoragerLogic) Storager(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
