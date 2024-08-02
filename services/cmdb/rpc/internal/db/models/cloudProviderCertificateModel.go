package models

import (
	"gorm.io/gorm"
)

type (
	CloudProviderCertificateModel interface {
		cloudProviderCertificateModel
		GetTableName() string
		WithTransaction(db *gorm.DB) CloudProviderCertificateModel
	}

	customCloudProviderCertificateModel struct {
		*defaultCloudProviderCertificateModel
	}
)

func (c *customCloudProviderCertificateModel) GetTableName() string {
	return c.tableName
}

func (c *customCloudProviderCertificateModel) WithTransaction(db *gorm.DB) CloudProviderCertificateModel {
	return &customCloudProviderCertificateModel{
		defaultCloudProviderCertificateModel: newCloudProviderCertificateModel(db),
	}
}

func NewCloudProviderCertificateModel(db *gorm.DB) CloudProviderCertificateModel {
	return &customCloudProviderCertificateModel{
		defaultCloudProviderCertificateModel: newCloudProviderCertificateModel(db),
	}
}
