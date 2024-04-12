package models

import (
	"gorm.io/gorm"
)

type WiredNetworkCard struct {
	gorm.Model
	Name      string  `json:"name" form:"name"`
	Price     float64 `json:"price" form:"price"`
	Interface string  `json:"interface" form:"interface"`
	Color     string  `json:"color" form:"color"`
}

func (WiredNetworkCard) TableName() string {
	return "wired_network_cards"
}
