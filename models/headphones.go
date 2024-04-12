package models

import (
	"gorm.io/gorm"
)

type Headphones struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price int `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 FrequencyResponse string `json:"frequency_response" form:"frequency_response"`
	 Microphone bool `json:"microphone" form:"microphone"`
	 Wireless bool `json:"wireless" form:"wireless"`
	 EnclosureType string `json:"enclosure_type" form:"enclosure_type"`
	 Color string `json:"color" form:"color"`
	
}

func (Headphones)TableName() string {
	return "headphones"
}
