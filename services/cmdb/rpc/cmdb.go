package main

import (
	"flag"
	"fmt"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/config"
	attributeServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/attribute"
	attributetypeServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/attributetype"
	categoryServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/category"
	modelServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/model"
	resourceServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resource"
	resourcegroupServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resourcegroup"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cmdb.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterResourceGroupServer(grpcServer, resourcegroupServer.NewResourceGroupServer(ctx))
		pb.RegisterCategoryServer(grpcServer, categoryServer.NewCategoryServer(ctx))
		pb.RegisterModelServer(grpcServer, modelServer.NewModelServer(ctx))
		pb.RegisterAttributeServer(grpcServer, attributeServer.NewAttributeServer(ctx))
		pb.RegisterAttributeTypeServer(grpcServer, attributetypeServer.NewAttributeTypeServer(ctx))
		pb.RegisterResourceServer(grpcServer, resourceServer.NewResourceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
