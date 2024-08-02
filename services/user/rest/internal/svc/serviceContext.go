package svc

import (
	"github.com/admgo/admgo/common/interceptor"
	commonmiddleware "github.com/admgo/admgo/common/middleware"
	"github.com/admgo/admgo/services/user/rest/internal/config"
	"github.com/admgo/admgo/services/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	CSRFTokenMiddleware      rest.Middleware
	AuthenticationMiddleware rest.Middleware
	UserRPC                  user.User
	Redis                    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.Rpc["user.rpc"], zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:                   c,
		UserRPC:                  user.NewUser(userRPC),
		AuthenticationMiddleware: commonmiddleware.NewAuthenticationMiddleware(c.Redis).Handle,
		CSRFTokenMiddleware:      commonmiddleware.NewCSRFTokenInterceptorMiddleware(c.Redis).Handle,
		Redis:                    redis.MustNewRedis(c.Redis, redis.WithPass(c.Redis.Pass)),
	}
}
