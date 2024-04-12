package models

import (
	"gorm.io/gorm"
)

type Case struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 Color string `json:"color" form:"color"`
	 Psu string `json:"psu" form:"psu"`
	 SidePanel string `json:"side_panel" form:"side_panel"`
	 ExternalVolume float64 `json:"external_volume" form:"external_volume"`
	 Internal35Bays int `json:"internal_35_bays" form:"internal_35_bays"`
	
}

func (Case)TableName() string {
	return "cases"
}
