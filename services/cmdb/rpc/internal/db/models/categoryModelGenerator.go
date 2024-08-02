package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	categoryModel interface {
		Insert(ctx context.Context, data *Category) error
		FindOne(ctx context.Context, id int64) (*Category, error)
		Update(ctx context.Context, data *Category) error
	}

	defaultCategoryModel struct {
		db        *gorm.DB
		tableName string
	}

	Category struct {
		BaseModel
		Name            string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		CategoryName    string `gorm:"column:category_name;type:varchar(30);not null;unique;comment:类别名(唯一)"`
		ResourceGroupID string `gorm:"column:resource_group_id;type:int unsigned;not null;comment:资源组ID"`
	}
)

func (emp Category) TableName() string {
	return TablePrefix + "_" + "categorys"
}

func newCategoryModel(conn *gorm.DB) *defaultCategoryModel {
	return &defaultCategoryModel{
		db:        conn,
		tableName: TablePrefix + "_" + "categorys",
	}
}

func (m *defaultCategoryModel) Insert(ctx context.Context, data *Category) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultCategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	var result Category
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultCategoryModel) Update(ctx context.Context, data *Category) error {
	return m.db.WithContext(ctx).Save(data).Error
}
