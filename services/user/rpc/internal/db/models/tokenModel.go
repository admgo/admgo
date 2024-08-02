package models

import (
	"gorm.io/gorm"
)

type (
	TokenModel interface {
		tokenModel
		GetTableName() string
		WithTransaction(db *gorm.DB) TokenModel
	}

	customTokenModel struct {
		*defaultTokenModel
	}
)

func (c *customTokenModel) GetTableName() string {
	return c.tableName
}

func (c *customTokenModel) WithTransaction(db *gorm.DB) TokenModel {
	return &customTokenModel{
		defaultTokenModel: newTokenModel(db),
	}
}

func NewTokenModel(db *gorm.DB) TokenModel {
	return &customTokenModel{
		defaultTokenModel: newTokenModel(db),
	}
}
