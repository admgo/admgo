package models

import (
	"gorm.io/gorm"
)

type (
	CloudProviderModel interface {
		cloudProviderModel
		GetTableName() string
		WithTransaction(db *gorm.DB) CloudProviderModel
	}

	customCloudProviderModel struct {
		*defaultCloudProviderModel
	}
)

func (c *customCloudProviderModel) GetTableName() string {
	return c.tableName
}

func (c *customCloudProviderModel) WithTransaction(db *gorm.DB) CloudProviderModel {
	return &customCloudProviderModel{
		defaultCloudProviderModel: newCloudProviderModel(db),
	}
}

func NewCloudProviderModel(db *gorm.DB) CloudProviderModel {
	return &customCloudProviderModel{
		defaultCloudProviderModel: newCloudProviderModel(db),
	}
}
