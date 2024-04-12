package models

import (
	"gorm.io/gorm"
)

type ThermalPaste struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Amount float64 `json:"amount" form:"amount"`
	
}

func (ThermalPaste)TableName() string {
	return "thermal_pastes"
}
