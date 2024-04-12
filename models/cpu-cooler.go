package models

import (
	"gorm.io/gorm"
)

type CpuCooler struct {
	gorm.Model
	Name       string  `json:"name" form:"name"`
	Price      float64 `json:"price" form:"price"`
	Rpm        int     `json:"rpm" form:"rpm"`
	NoiseLevel float64 `json:"noise_level" form:"noise_level"`
	Color      string  `json:"color" form:"color"`
	Size       string  `json:"size" form:"size"`
}

func (CpuCooler) TableName() string {
	return "cpu_coolers"
}
