package svc

import (
	"github.com/admgo/admgo/common/interceptor"
	commonmiddleware "github.com/admgo/admgo/common/middleware"
	"github.com/admgo/admgo/pkg/db"
	dbconfig "github.com/admgo/admgo/pkg/db/config"
	"github.com/admgo/admgo/services/auth/api/internal/config"
	"github.com/admgo/admgo/services/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	CSRFTokenMiddleware      rest.Middleware
	AuthenticationMiddleware rest.Middleware
	DB                       *db.DB
	UserRPC                  user.User
	Redis                    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := db.MustNewMysql(&dbconfig.Config{
		Username:     c.DB.Username,
		Password:     c.DB.Password,
		Host:         c.DB.Host,
		Port:         c.DB.Port,
		Database:     c.DB.Database,
		Charset:      c.DB.Charset,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
		TimeOut:      c.DB.TimeOut,
		ReadTimeOut:  c.DB.ReadTimeOut,
		WriteTimeOut: c.DB.WriteTimeOut,
	})
	userRPC := zrpc.MustNewClient(c.Rpc["user.rpc"], zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:                   c,
		DB:                       db,
		UserRPC:                  user.NewUser(userRPC),
		AuthenticationMiddleware: commonmiddleware.NewAuthenticationMiddleware(c.Redis).Handle,
		CSRFTokenMiddleware:      commonmiddleware.NewCSRFTokenInterceptorMiddleware(c.Redis).Handle,
		Redis:                    redis.MustNewRedis(c.Redis, redis.WithPass(c.Redis.Pass)),
	}
}
