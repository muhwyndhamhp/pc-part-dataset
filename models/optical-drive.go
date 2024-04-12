package models

import (
	"gorm.io/gorm"
)

type OpticalDrive struct {
	gorm.Model
	Name     string  `json:"name" form:"name"`
	Price    float64 `json:"price" form:"price"`
	Bd       int     `json:"bd" form:"bd"`
	Dvd      int     `json:"dvd" form:"dvd"`
	Cd       int     `json:"cd" form:"cd"`
	BdWrite  string  `json:"bd_write" form:"bd_write"`
	DvdWrite string  `json:"dvd_write" form:"dvd_write"`
	CdWrite  string  `json:"cd_write" form:"cd_write"`
}

func (OpticalDrive) TableName() string {
	return "optical_drives"
}
