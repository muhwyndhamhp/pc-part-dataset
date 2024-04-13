package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed thermal-paste.csv
var ThermalPasteCsv []byte

type ThermalPaste struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Amount float64 `json:"amount" form:"amount"`
	
}

func (ThermalPaste)TableName() string {
	return "thermal_pastes"
}

func (ThermalPaste)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(ThermalPasteCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*ThermalPaste{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := ThermalPaste{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Amount: utils.ToFloat64(records[i][2]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
