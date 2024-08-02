package main

import (
	"flag"
	"github.com/admgo/admgo/pkg/db"
	dbconfig "github.com/admgo/admgo/pkg/db/config"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/config"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/db/models"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var configFile = flag.String("f", "../../etc/cmdb.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	dbInstance := db.MustNewMysql(&dbconfig.Config{
		Username:     c.DB.Username,
		Password:     c.DB.Password,
		Host:         c.DB.Host,
		Port:         c.DB.Port,
		Database:     c.DB.Database,
		Charset:      c.DB.Charset,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxLifetime:  c.DB.MaxLifetime,
		TimeOut:      c.DB.TimeOut,
		WriteTimeOut: c.DB.WriteTimeOut,
		ReadTimeOut:  c.DB.ReadTimeOut,
	})
	err := dbInstance.AutoMigrate(
		&models.Model{},
		&models.Category{},
		&models.ResourceGroup{},
		&models.CloudProvider{},
		&models.CloudProviderCertificate{},
	)
	if err != nil {
		panic(err)
	}
}
