package models

import (
	"gorm.io/gorm"
)

type CaseAccessory struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price string `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 FormFactor float64 `json:"form_factor" form:"form_factor"`
	
}

func (CaseAccessory)TableName() string {
	return "case_accessories"
}
