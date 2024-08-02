package code

import "github.com/admgo/admgo/pkg/errorx"

var (
	UserEmpty                   = errorx.NewRpcStatus(120101, "no user")
	IncorrectUserNameORPassword = errorx.NewRpcStatus(120102, "incorrect user name or password")
	MysqlErr                    = errorx.NewRpcStatus(120102, "internal error")
)
