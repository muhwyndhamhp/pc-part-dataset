package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed ups.csv
var UPSCsv []byte

type UPS struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 CapacityW float64 `json:"capacity_w" form:"capacity_w"`
	 CapacityVa float64 `json:"capacity_va" form:"capacity_va"`
	
}

func (UPS)TableName() string {
	return "ups"
}

func (UPS)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(UPSCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*UPS{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := UPS{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			CapacityW: utils.ToFloat64(records[i][2]),
			CapacityVa: utils.ToFloat64(records[i][3]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
