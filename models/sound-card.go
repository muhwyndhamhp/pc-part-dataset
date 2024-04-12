package models

import (
	"gorm.io/gorm"
)

type SoundCard struct {
	gorm.Model
	Name         string  `json:"name" form:"name"`
	Price        float64 `json:"price" form:"price"`
	Channels     float64 `json:"channels" form:"channels"`
	DigitalAudio int     `json:"digital_audio" form:"digital_audio"`
	Snr          int     `json:"snr" form:"snr"`
	SampleRate   int     `json:"sample_rate" form:"sample_rate"`
	Chipset      string  `json:"chipset" form:"chipset"`
	Interface    string  `json:"interface" form:"interface"`
}

func (SoundCard) TableName() string {
	return "sound_cards"
}
