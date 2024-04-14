package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed cpu-cooler.csv
var CpuCoolerCsv []byte

type CpuCooler struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Rpm float64 `json:"rpm" form:"rpm"`
	 NoiseLevel float64 `json:"noise_level" form:"noise_level"`
	 Color string `json:"color" form:"color"`
	 Size string `json:"size" form:"size"`
	
}

func (CpuCooler)TableName() string {
	return "cpu_coolers"
}

func (CpuCooler)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(CpuCoolerCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*CpuCooler{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := CpuCooler{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Rpm: utils.ToFloat64(records[i][2]),
			NoiseLevel: utils.ToFloat64(records[i][3]),
			Color: records[i][4],
			Size: records[i][5],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
