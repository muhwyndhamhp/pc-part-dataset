package models

import (
	"gorm.io/gorm"
)

type FanController struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Channels int `json:"channels" form:"channels"`
	 ChannelWattage string `json:"channel_wattage" form:"channel_wattage"`
	 Pwm bool `json:"pwm" form:"pwm"`
	 FormFactor string `json:"form_factor" form:"form_factor"`
	 Color string `json:"color" form:"color"`
	
}

func (FanController)TableName() string {
	return "fan_controllers"
}
