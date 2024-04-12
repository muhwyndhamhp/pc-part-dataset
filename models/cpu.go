package models

import (
	"gorm.io/gorm"
)

type CPU struct {
	gorm.Model
	Name       string  `json:"name" form:"name"`
	Price      int     `json:"price" form:"price"`
	CoreCount  int     `json:"core_count" form:"core_count"`
	CoreClock  float64 `json:"core_clock" form:"core_clock"`
	BoostClock int     `json:"boost_clock" form:"boost_clock"`
	Tdp        int     `json:"tdp" form:"tdp"`
	Graphics   string  `json:"graphics" form:"graphics"`
	Smt        bool    `json:"smt" form:"smt"`
}

func (CPU) TableName() string {
	return "cpus"
}
