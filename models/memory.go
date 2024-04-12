package models

import (
	"gorm.io/gorm"
)

type Memory struct {
	gorm.Model
	Name             string  `json:"name" form:"name"`
	Price            float64 `json:"price" form:"price"`
	Speed            string  `json:"speed" form:"speed"`
	Modules          string  `json:"modules" form:"modules"`
	PricePerGb       float64 `json:"price_per_gb" form:"price_per_gb"`
	Color            string  `json:"color" form:"color"`
	FirstWordLatency int     `json:"first_word_latency" form:"first_word_latency"`
	CasLatency       int     `json:"cas_latency" form:"cas_latency"`
}

func (Memory) TableName() string {
	return "memories"
}
