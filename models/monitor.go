package models

import (
	"gorm.io/gorm"
)

type Monitor struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 ScreenSize int `json:"screen_size" form:"screen_size"`
	 Resolution string `json:"resolution" form:"resolution"`
	 RefreshRate int `json:"refresh_rate" form:"refresh_rate"`
	 ResponseTime float64 `json:"response_time" form:"response_time"`
	 PanelType string `json:"panel_type" form:"panel_type"`
	 AspectRatio string `json:"aspect_ratio" form:"aspect_ratio"`
	
}

func (Monitor)TableName() string {
	return "monitors"
}
