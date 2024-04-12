package models

import (
	"gorm.io/gorm"
)

type VideoCard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Chipset string `json:"chipset" form:"chipset"`
	 Memory int `json:"memory" form:"memory"`
	 CoreClock int `json:"core_clock" form:"core_clock"`
	 BoostClock int `json:"boost_clock" form:"boost_clock"`
	 Color string `json:"color" form:"color"`
	 Length int `json:"length" form:"length"`
	
}

func (VideoCard)TableName() string {
	return "video_cards"
}
