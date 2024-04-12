package models

import (
	"gorm.io/gorm"
)

type UPS struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price int `json:"price" form:"price"`
	 CapacityW int `json:"capacity_w" form:"capacity_w"`
	 CapacityVa int `json:"capacity_va" form:"capacity_va"`
	
}

func (UPS)TableName() string {
	return "ups"
}
