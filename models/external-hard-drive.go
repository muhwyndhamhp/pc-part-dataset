package models

import (
	"gorm.io/gorm"
)

type ExternalHardDrive struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price int `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 Interface string `json:"interface" form:"interface"`
	 Capacity int `json:"capacity" form:"capacity"`
	 PricePerGb float64 `json:"price_per_gb" form:"price_per_gb"`
	 Color string `json:"color" form:"color"`
	
}

func (ExternalHardDrive)TableName() string {
	return "external_hard_drives"
}
