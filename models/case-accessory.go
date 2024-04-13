package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed case-accessory.csv
var CaseAccessoryCsv []byte

type CaseAccessory struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price string `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 FormFactor float64 `json:"form_factor" form:"form_factor"`
	
}

func (CaseAccessory)TableName() string {
	return "case_accessories"
}

func (CaseAccessory)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(CaseAccessoryCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*CaseAccessory{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := CaseAccessory{
			Name: records[i][0],
			Price: records[i][1],
			Type: records[i][2],
			FormFactor: utils.ToFloat64(records[i][3]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
