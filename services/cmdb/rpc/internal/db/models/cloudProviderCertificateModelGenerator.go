package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	cloudProviderCertificateModel interface {
		Insert(ctx context.Context, data *CloudProviderCertificate) error
		FindOne(ctx context.Context, id int64) (*CloudProviderCertificate, error)
		Update(ctx context.Context, data *CloudProviderCertificate) error
	}

	defaultCloudProviderCertificateModel struct {
		db        *gorm.DB
		tableName string
	}

	CloudProviderCertificate struct {
		BaseModel
		Name                         string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		CloudProviderCertificateName string `gorm:"column:cloud_provider_certificate_name;type:varchar(30);not null;unique;comment:云厂商凭证(唯一)"`
		CloudProviderID              string `gorm:"column:cloud_provider_id;type:int unsigned;not null;comment:云厂商ID"`
		AccessKeyID                  string `gorm:"column:access_key_id;type:varchar(100);not null;comment:AK"`
		SecretAccessKey              string `gorm:"column:secret_access_key;type:varchar(100);not null;comment:SK"`
	}
)

func (emp CloudProviderCertificate) TableName() string {
	return TablePrefix + "_" + "cloud_provider_certificates"
}

func newCloudProviderCertificateModel(conn *gorm.DB) *defaultCloudProviderCertificateModel {
	return &defaultCloudProviderCertificateModel{
		db:        conn,
		tableName: TablePrefix + "_" + "cloud_provider_certificates",
	}
}

func (m *defaultCloudProviderCertificateModel) Insert(ctx context.Context, data *CloudProviderCertificate) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultCloudProviderCertificateModel) FindOne(ctx context.Context, id int64) (*CloudProviderCertificate, error) {
	var result CloudProviderCertificate
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultCloudProviderCertificateModel) Update(ctx context.Context, data *CloudProviderCertificate) error {
	return m.db.WithContext(ctx).Save(data).Error
}
