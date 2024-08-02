package models

import (
	"gorm.io/gorm"
)

type (
	CategoryModel interface {
		categoryModel
		GetTableName() string
		WithTransaction(db *gorm.DB) CategoryModel
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

func (c *customCategoryModel) GetTableName() string {
	return c.tableName
}

func (c *customCategoryModel) WithTransaction(db *gorm.DB) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(db),
	}
}

func NewCategoryModel(db *gorm.DB) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(db),
	}
}
