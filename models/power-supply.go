package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed power-supply.csv
var PowerSupplyCsv []byte

type PowerSupply struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 Efficiency string `json:"efficiency" form:"efficiency"`
	 Wattage int `json:"wattage" form:"wattage"`
	 Modular string `json:"modular" form:"modular"`
	 Color string `json:"color" form:"color"`
	
}

func (PowerSupply)TableName() string {
	return "power_supplies"
}

func (PowerSupply)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(PowerSupplyCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*PowerSupply{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := PowerSupply{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Type: records[i][2],
			Efficiency: records[i][3],
			Wattage: utils.ToInt(records[i][4]),
			Modular: records[i][5],
			Color: records[i][6],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
