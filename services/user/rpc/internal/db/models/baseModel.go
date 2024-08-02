package models

import (
	"time"
)

type (
	BaseModel struct {
		ID        uint32    `gorm:"primaryKey;autoIncrement;comment:主键ID"`
		CreatedAt time.Time `gorm:"column: created_at;comment:创建时间"`
		UpdatedAt time.Time `gorm:"column: updated_at;comment:更新时间"`
	}
)

const TablePrefix = "user"
