package modellogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResourceModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateResourceModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceModelLogic {
	return &CreateResourceModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateResourceModel 创建资源模型
func (l *CreateResourceModelLogic) CreateResourceModel(in *pb.CreateModelRequest) (*pb.CreateModelResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateModelResponse{}, nil
}
