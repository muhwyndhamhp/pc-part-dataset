package models

import (
	"gorm.io/gorm"
)

type Speakers struct {
	gorm.Model
	Name              string  `json:"name" form:"name"`
	Price             float64 `json:"price" form:"price"`
	Configuration     int     `json:"configuration" form:"configuration"`
	Wattage           int     `json:"wattage" form:"wattage"`
	FrequencyResponse string  `json:"frequency_response" form:"frequency_response"`
	Color             string  `json:"color" form:"color"`
}

func (Speakers) TableName() string {
	return "speakers"
}
