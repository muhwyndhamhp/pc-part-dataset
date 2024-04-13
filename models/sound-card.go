package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed sound-card.csv
var SoundCardCsv []byte

type SoundCard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Channels float64 `json:"channels" form:"channels"`
	 DigitalAudio int `json:"digital_audio" form:"digital_audio"`
	 Snr int `json:"snr" form:"snr"`
	 SampleRate int `json:"sample_rate" form:"sample_rate"`
	 Chipset string `json:"chipset" form:"chipset"`
	 Interface string `json:"interface" form:"interface"`
	
}

func (SoundCard)TableName() string {
	return "sound_cards"
}

func (SoundCard)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(SoundCardCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*SoundCard{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := SoundCard{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Channels: utils.ToFloat64(records[i][2]),
			DigitalAudio: utils.ToInt(records[i][3]),
			Snr: utils.ToInt(records[i][4]),
			SampleRate: utils.ToInt(records[i][5]),
			Chipset: records[i][6],
			Interface: records[i][7],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
