package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed cpu.csv
var CPUCsv []byte

type CPU struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 CoreCount float64 `json:"core_count" form:"core_count"`
	 CoreClock float64 `json:"core_clock" form:"core_clock"`
	 BoostClock float64 `json:"boost_clock" form:"boost_clock"`
	 Tdp float64 `json:"tdp" form:"tdp"`
	 Graphics string `json:"graphics" form:"graphics"`
	 Smt bool `json:"smt" form:"smt"`
	
}

func (CPU)TableName() string {
	return "cpus"
}

func (CPU)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(CPUCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*CPU{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := CPU{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			CoreCount: utils.ToFloat64(records[i][2]),
			CoreClock: utils.ToFloat64(records[i][3]),
			BoostClock: utils.ToFloat64(records[i][4]),
			Tdp: utils.ToFloat64(records[i][5]),
			Graphics: records[i][6],
			Smt: utils.ToBool(records[i][7]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
