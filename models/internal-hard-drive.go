package models

import (
	"gorm.io/gorm"
)

type InternalHardDrive struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Capacity int `json:"capacity" form:"capacity"`
	 PricePerGb float64 `json:"price_per_gb" form:"price_per_gb"`
	 Type string `json:"type" form:"type"`
	 Cache int `json:"cache" form:"cache"`
	 FormFactor string `json:"form_factor" form:"form_factor"`
	 Interface string `json:"interface" form:"interface"`
	
}

func (InternalHardDrive)TableName() string {
	return "internal_hard_drives"
}
