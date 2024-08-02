package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	modelModel interface {
		Insert(ctx context.Context, data *Model) error
		FindOne(ctx context.Context, id int64) (*Model, error)
		Update(ctx context.Context, data *Model) error
	}

	defaultModelModel struct {
		db        *gorm.DB
		tableName string
	}

	Model struct {
		BaseModel
		Name       string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		ModelName  string `gorm:"column:model_name;type:varchar(30);not null;unique;comment:模型名(唯一)"`
		CategoryID string `gorm:"column:category_id;type:int unsigned;not null;comment:类别ID"`
	}
)

func (emp Model) TableName() string {
	return TablePrefix + "_" + "models"
}

func newModelModel(conn *gorm.DB) *defaultModelModel {
	return &defaultModelModel{
		db:        conn,
		tableName: TablePrefix + "_" + "models",
	}
}

func (m *defaultModelModel) Insert(ctx context.Context, data *Model) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultModelModel) FindOne(ctx context.Context, id int64) (*Model, error) {
	var result Model
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultModelModel) Update(ctx context.Context, data *Model) error {
	return m.db.WithContext(ctx).Save(data).Error
}
