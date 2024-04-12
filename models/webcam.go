package models

import (
	"gorm.io/gorm"
)

type Webcam struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Resolutions string `json:"resolutions" form:"resolutions"`
	 Connection string `json:"connection" form:"connection"`
	 FocusType string `json:"focus_type" form:"focus_type"`
	 Os string `json:"os" form:"os"`
	 Fov float64 `json:"fov" form:"fov"`
	
}

func (Webcam)TableName() string {
	return "webcams"
}
