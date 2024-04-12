package models

import (
	"gorm.io/gorm"
)

type Mouse struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 TrackingMethod string `json:"tracking_method" form:"tracking_method"`
	 ConnectionType string `json:"connection_type" form:"connection_type"`
	 MaxDpi int `json:"max_dpi" form:"max_dpi"`
	 HandOrientation string `json:"hand_orientation" form:"hand_orientation"`
	 Color string `json:"color" form:"color"`
	
}

func (Mouse)TableName() string {
	return "mouses"
}
