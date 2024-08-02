package models

import (
	"gorm.io/gorm"
)

type (
	ResourceGroupModel interface {
		resourceGroupModel
		GetTableName() string
		WithTransaction(db *gorm.DB) ResourceGroupModel
	}

	customResourceGroupModel struct {
		*defaultResourceGroupModel
	}
)

func (c *customResourceGroupModel) GetTableName() string {
	return c.tableName
}

func (c *customResourceGroupModel) WithTransaction(db *gorm.DB) ResourceGroupModel {
	return &customResourceGroupModel{
		defaultResourceGroupModel: newResourceGroupModel(db),
	}
}

func NewResourceGroupModel(db *gorm.DB) ResourceGroupModel {
	return &customResourceGroupModel{
		defaultResourceGroupModel: newResourceGroupModel(db),
	}
}
