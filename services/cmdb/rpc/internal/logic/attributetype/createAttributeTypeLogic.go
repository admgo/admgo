package attributetypelogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAttributeTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAttributeTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAttributeTypeLogic {
	return &CreateAttributeTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateAttributeType 创建资源属性类型
func (l *CreateAttributeTypeLogic) CreateAttributeType(in *pb.CreateAttributeTypeRequest) (*pb.CreateAttributeTypeResponse, error) {

	return &pb.CreateAttributeTypeResponse{}, nil
}
