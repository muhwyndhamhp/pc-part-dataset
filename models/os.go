package models

import (
	"gorm.io/gorm"
)

type OS struct {
	gorm.Model
	Name      string  `json:"name" form:"name"`
	Price     float64 `json:"price" form:"price"`
	Mode      int     `json:"mode" form:"mode"`
	MaxMemory int     `json:"max_memory" form:"max_memory"`
}

func (OS) TableName() string {
	return "os"
}
