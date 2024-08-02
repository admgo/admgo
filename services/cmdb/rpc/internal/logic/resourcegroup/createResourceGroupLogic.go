package resourcegrouplogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResourceGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateResourceGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceGroupLogic {
	return &CreateResourceGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateResourceGroup 创建资源组
func (l *CreateResourceGroupLogic) CreateResourceGroup(in *pb.CreateResourceGroupRequest) (*pb.CreateResourceGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateResourceGroupResponse{}, nil
}
