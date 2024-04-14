package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed memory.csv
var MemoryCsv []byte

type Memory struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Speed string `json:"speed" form:"speed"`
	 Modules string `json:"modules" form:"modules"`
	 PricePerGb float64 `json:"price_per_gb" form:"price_per_gb"`
	 Color string `json:"color" form:"color"`
	 FirstWordLatency float64 `json:"first_word_latency" form:"first_word_latency"`
	 CasLatency float64 `json:"cas_latency" form:"cas_latency"`
	
}

func (Memory)TableName() string {
	return "memories"
}

func (Memory)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(MemoryCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Memory{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Memory{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Speed: records[i][2],
			Modules: records[i][3],
			PricePerGb: utils.ToFloat64(records[i][4]),
			Color: records[i][5],
			FirstWordLatency: utils.ToFloat64(records[i][6]),
			CasLatency: utils.ToFloat64(records[i][7]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
