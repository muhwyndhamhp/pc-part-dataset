package models

import (
	"gorm.io/gorm"
)

type Keyboard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Style string `json:"style" form:"style"`
	 Switches string `json:"switches" form:"switches"`
	 Backlit string `json:"backlit" form:"backlit"`
	 Tenkeyless bool `json:"tenkeyless" form:"tenkeyless"`
	 ConnectionType string `json:"connection_type" form:"connection_type"`
	 Color string `json:"color" form:"color"`
	
}

func (Keyboard)TableName() string {
	return "keyboards"
}
