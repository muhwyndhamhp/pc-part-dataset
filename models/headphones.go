package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed headphones.csv
var HeadphonesCsv []byte

type Headphones struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 FrequencyResponse string `json:"frequency_response" form:"frequency_response"`
	 Microphone bool `json:"microphone" form:"microphone"`
	 Wireless bool `json:"wireless" form:"wireless"`
	 EnclosureType string `json:"enclosure_type" form:"enclosure_type"`
	 Color string `json:"color" form:"color"`
	
}

func (Headphones)TableName() string {
	return "headphones"
}

func (Headphones)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(HeadphonesCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Headphones{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Headphones{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Type: records[i][2],
			FrequencyResponse: records[i][3],
			Microphone: utils.ToBool(records[i][4]),
			Wireless: utils.ToBool(records[i][5]),
			EnclosureType: records[i][6],
			Color: records[i][7],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
