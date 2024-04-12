package models

import (
	"gorm.io/gorm"
)

type WirelessNetworkCard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Protocol string `json:"protocol" form:"protocol"`
	 Interface string `json:"interface" form:"interface"`
	 Color string `json:"color" form:"color"`
	
}

func (WirelessNetworkCard)TableName() string {
	return "wireless_network_cards"
}
