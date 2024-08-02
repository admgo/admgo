package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	tokenModel interface {
		Insert(ctx context.Context, data *Token) error
		FindOne(ctx context.Context, id int64) (*Token, error)
		Update(ctx context.Context, data *Token) error
	}

	defaultTokenModel struct {
		db        *gorm.DB
		tableName string
	}

	Token struct {
		BaseModel
		Name           string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		UserName       string `gorm:"column:username;type:varchar(30);not null;unique;comment:用户名(唯一)"`
		Password       string `gorm:"column:password;type:varchar(100);not null;comment:密码"`
		Email          string `gorm:"column:email;type:varchar(60);comment:邮箱"`
		EmployeeNumber string `gorm:"column:employee_number;type:varchar(30);comment:员工号"`
		Phone          string `gorm:"column:phone;type:varchar(30);comment:手机号"`
		Avatar         string `gorm:"column:avatar;type:varchar(255);comment:头像"`
		IsAdmin        bool   `gorm:"column:is_admin;type:bool;not null;comment:是否是管理员"`
	}
)

func (emp Token) TableName() string {
	return TablePrefix + "_" + "tokens"
}

func newTokenModel(conn *gorm.DB) *defaultTokenModel {
	return &defaultTokenModel{
		db:        conn,
		tableName: TablePrefix + "_" + "tokens",
	}
}

func (m *defaultTokenModel) Insert(ctx context.Context, data *Token) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultTokenModel) FindOne(ctx context.Context, id int64) (*Token, error) {
	var result Token
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultTokenModel) Update(ctx context.Context, data *Token) error {
	return m.db.WithContext(ctx).Save(data).Error
}
