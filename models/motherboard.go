package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed motherboard.csv
var MotherboardCsv []byte

type Motherboard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price int `json:"price" form:"price"`
	 Socket string `json:"socket" form:"socket"`
	 FormFactor string `json:"form_factor" form:"form_factor"`
	 MaxMemory int `json:"max_memory" form:"max_memory"`
	 MemorySlots int `json:"memory_slots" form:"memory_slots"`
	 Color string `json:"color" form:"color"`
	
}

func (Motherboard)TableName() string {
	return "motherboards"
}

func (Motherboard)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(MotherboardCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Motherboard{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Motherboard{
			Name: records[i][0],
			Price: utils.ToInt(records[i][1]),
			Socket: records[i][2],
			FormFactor: records[i][3],
			MaxMemory: utils.ToInt(records[i][4]),
			MemorySlots: utils.ToInt(records[i][5]),
			Color: records[i][6],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
