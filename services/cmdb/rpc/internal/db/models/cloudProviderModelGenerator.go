package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	cloudProviderModel interface {
		Insert(ctx context.Context, data *CloudProvider) error
		FindOne(ctx context.Context, id int64) (*CloudProvider, error)
		Update(ctx context.Context, data *CloudProvider) error
	}

	defaultCloudProviderModel struct {
		db        *gorm.DB
		tableName string
	}

	CloudProvider struct {
		BaseModel
		Name              string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		CloudProviderName string `gorm:"column:cloud_provider_name;type:varchar(30);not null;unique;comment:云厂商名(唯一)"`
	}
)

func (emp CloudProvider) TableName() string {
	return TablePrefix + "_" + "cloud_providers"
}

func newCloudProviderModel(conn *gorm.DB) *defaultCloudProviderModel {
	return &defaultCloudProviderModel{
		db:        conn,
		tableName: TablePrefix + "_" + "cloud_providers",
	}
}

func (m *defaultCloudProviderModel) Insert(ctx context.Context, data *CloudProvider) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultCloudProviderModel) FindOne(ctx context.Context, id int64) (*CloudProvider, error) {
	var result CloudProvider
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultCloudProviderModel) Update(ctx context.Context, data *CloudProvider) error {
	return m.db.WithContext(ctx).Save(data).Error
}
