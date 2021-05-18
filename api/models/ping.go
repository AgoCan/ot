package models

import (
	"gorm.io/gorm"
)

// Ping 测试
type Ping struct {
	gorm.Model
	Msg string `db:"msg" json:"msg"`
}
