package models

import (
	"gorm.io/gorm"
)

type CaseFan struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Size int `json:"size" form:"size"`
	 Color string `json:"color" form:"color"`
	 Rpm string `json:"rpm" form:"rpm"`
	 Airflow string `json:"airflow" form:"airflow"`
	 NoiseLevel string `json:"noise_level" form:"noise_level"`
	 Pwm bool `json:"pwm" form:"pwm"`
	
}

func (CaseFan)TableName() string {
	return "case_fans"
}
