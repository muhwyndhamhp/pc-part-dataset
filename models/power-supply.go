package models

import (
	"gorm.io/gorm"
)

type PowerSupply struct {
	gorm.Model
	Name       string  `json:"name" form:"name"`
	Price      float64 `json:"price" form:"price"`
	Type       string  `json:"type" form:"type"`
	Efficiency string  `json:"efficiency" form:"efficiency"`
	Wattage    int     `json:"wattage" form:"wattage"`
	Modular    string  `json:"modular" form:"modular"`
	Color      string  `json:"color" form:"color"`
}

func (PowerSupply) TableName() string {
	return "power_supplies"
}
