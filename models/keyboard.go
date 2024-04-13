package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed keyboard.csv
var KeyboardCsv []byte

type Keyboard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Style string `json:"style" form:"style"`
	 Switches string `json:"switches" form:"switches"`
	 Backlit string `json:"backlit" form:"backlit"`
	 Tenkeyless bool `json:"tenkeyless" form:"tenkeyless"`
	 ConnectionType string `json:"connection_type" form:"connection_type"`
	 Color string `json:"color" form:"color"`
	
}

func (Keyboard)TableName() string {
	return "keyboards"
}

func (Keyboard)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(KeyboardCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Keyboard{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Keyboard{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Style: records[i][2],
			Switches: records[i][3],
			Backlit: records[i][4],
			Tenkeyless: utils.ToBool(records[i][5]),
			ConnectionType: records[i][6],
			Color: records[i][7],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
