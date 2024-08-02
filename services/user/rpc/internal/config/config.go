package config

import (
	db "github.com/admgo/admgo/pkg/db/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB db.Config
}
