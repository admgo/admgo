package main

import (
	"flag"
	"github.com/admgo/admgo/common/interceptor"
	"github.com/admgo/admgo/services/user/rpc/pb"
	"github.com/admgo/admgo/utils/tip"

	"github.com/admgo/admgo/services/user/rpc/internal/config"
	"github.com/admgo/admgo/services/user/rpc/internal/server"
	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(interceptor.ServerErrorInterceptor())
	tip.Tip(&c)
	s.Start()
}
