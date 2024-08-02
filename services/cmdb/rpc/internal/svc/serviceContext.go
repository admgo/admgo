package svc

import (
	"github.com/admgo/admgo/pkg/db"
	dbconfig "github.com/admgo/admgo/pkg/db/config"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/config"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/db/models"
)

type ServiceContext struct {
	Config        config.Config
	DB            *db.DB
	ResourceGroup models.ResourceGroupModel
	Model         models.ModelModel
	Category      models.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
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
	return &ServiceContext{
		Config:        c,
		DB:            dbInstance,
		ResourceGroup: models.NewResourceGroupModel(dbInstance.DB),
		Model:         models.NewModelModel(dbInstance.DB),
		Category:      models.NewCategoryModel(dbInstance.DB),
	}
}
