package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed case-fan.csv
var CaseFanCsv []byte

type CaseFan struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Size float64 `json:"size" form:"size"`
	 Color string `json:"color" form:"color"`
	 Rpm string `json:"rpm" form:"rpm"`
	 Airflow string `json:"airflow" form:"airflow"`
	 NoiseLevel string `json:"noise_level" form:"noise_level"`
	 Pwm bool `json:"pwm" form:"pwm"`
	
}

func (CaseFan)TableName() string {
	return "case_fans"
}

func (CaseFan)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(CaseFanCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*CaseFan{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := CaseFan{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Size: utils.ToFloat64(records[i][2]),
			Color: records[i][3],
			Rpm: records[i][4],
			Airflow: records[i][5],
			NoiseLevel: records[i][6],
			Pwm: utils.ToBool(records[i][7]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
