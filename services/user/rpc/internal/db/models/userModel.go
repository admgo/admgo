package models

import (
	"fmt"
	"gorm.io/gorm"
)

type (
	UserModel interface {
		userModel
		GetTableName() string
		WithTransaction(db *gorm.DB) UserModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (c *customUserModel) GetTableName() string {
	return c.tableName
}

func (c *customUserModel) WithTransaction(db *gorm.DB) UserModel {
	fmt.Println(&customUserModel{
		defaultUserModel: newUserModel(db),
	})
	return &customUserModel{
		defaultUserModel: newUserModel(db),
	}
}

func NewUserModel(conn *gorm.DB) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
