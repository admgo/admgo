package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	resourceGroupModel interface {
		Insert(ctx context.Context, data *ResourceGroup) error
		FindOne(ctx context.Context, id int64) (*ResourceGroup, error)
		Update(ctx context.Context, data *ResourceGroup) error
	}

	defaultResourceGroupModel struct {
		db        *gorm.DB
		tableName string
	}

	ResourceGroup struct {
		BaseModel
		Name              string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		ResourceGroupName string `gorm:"column:resource_group_name;type:varchar(30);not null;unique;comment:资源组名(唯一)"`
	}
)

func (emp ResourceGroup) TableName() string {
	return TablePrefix + "_" + "resource_groups"
}

func newResourceGroupModel(conn *gorm.DB) *defaultResourceGroupModel {
	return &defaultResourceGroupModel{
		db:        conn,
		tableName: TablePrefix + "_" + "resource_groups",
	}
}

func (m *defaultResourceGroupModel) Insert(ctx context.Context, data *ResourceGroup) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultResourceGroupModel) FindOne(ctx context.Context, id int64) (*ResourceGroup, error) {
	var result ResourceGroup
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultResourceGroupModel) Update(ctx context.Context, data *ResourceGroup) error {
	return m.db.WithContext(ctx).Save(data).Error
}
