package models

import (
	"gorm.io/gorm"
)

type (
	ModelModel interface {
		modelModel
		GetTableName() string
		WithTransaction(db *gorm.DB) ModelModel
	}

	customModelModel struct {
		*defaultModelModel
	}
)

func (c *customModelModel) GetTableName() string {
	return c.tableName
}

func (c *customModelModel) WithTransaction(db *gorm.DB) ModelModel {
	return &customModelModel{
		defaultModelModel: newModelModel(db),
	}
}

func NewModelModel(db *gorm.DB) ModelModel {
	return &customModelModel{
		defaultModelModel: newModelModel(db),
	}
}
