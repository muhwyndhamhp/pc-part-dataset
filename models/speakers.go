package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed speakers.csv
var SpeakersCsv []byte

type Speakers struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Configuration float64 `json:"configuration" form:"configuration"`
	 Wattage float64 `json:"wattage" form:"wattage"`
	 FrequencyResponse string `json:"frequency_response" form:"frequency_response"`
	 Color string `json:"color" form:"color"`
	
}

func (Speakers)TableName() string {
	return "speakers"
}

func (Speakers)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(SpeakersCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Speakers{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Speakers{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Configuration: utils.ToFloat64(records[i][2]),
			Wattage: utils.ToFloat64(records[i][3]),
			FrequencyResponse: records[i][4],
			Color: records[i][5],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
