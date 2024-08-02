package main

import (
	"flag"
	"fmt"
	"github.com/admgo/admgo/pkg/errorx"
	"github.com/admgo/admgo/pkg/response"
	"github.com/admgo/admgo/pkg/validator"
	"github.com/admgo/admgo/utils/tip"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/admgo/admgo/services/user/rest/internal/config"
	"github.com/admgo/admgo/services/user/rest/internal/handler"
	"github.com/admgo/admgo/services/user/rest/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	//server.Use(middleware.NewCORSMiddleware(&middleware.CORSMiddlewareOptions{}).Handle)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(errorx.ErrHandler)
	httpx.SetValidator(validator.NewValidate())
	httpx.SetOkHandler(response.OKResponse)

	tip.Tip(c)
	fmt.Printf("Starting [%s] service at %s:%d...\n", c.Name, c.Host, c.Port)
	server.Start()
}
