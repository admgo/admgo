package config

import (
	db "github.com/admgo/admgo/pkg/db/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB    db.Config
	Rpc   map[string]zrpc.RpcClientConf
	Redis redis.RedisConf
}
