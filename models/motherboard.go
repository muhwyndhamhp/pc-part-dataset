package models

import (
	"gorm.io/gorm"
)

type Motherboard struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Socket      string `json:"socket" form:"socket"`
	FormFactor  string `json:"form_factor" form:"form_factor"`
	MaxMemory   int    `json:"max_memory" form:"max_memory"`
	MemorySlots int    `json:"memory_slots" form:"memory_slots"`
	Color       string `json:"color" form:"color"`
}

func (Motherboard) TableName() string {
	return "motherboards"
}
