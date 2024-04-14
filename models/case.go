package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed case.csv
var CaseCsv []byte

type Case struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 Color string `json:"color" form:"color"`
	 Psu string `json:"psu" form:"psu"`
	 SidePanel string `json:"side_panel" form:"side_panel"`
	 ExternalVolume float64 `json:"external_volume" form:"external_volume"`
	 Internal35Bays float64 `json:"internal_35_bays" form:"internal_35_bays"`
	
}

func (Case)TableName() string {
	return "cases"
}

func (Case)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(CaseCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Case{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Case{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Type: records[i][2],
			Color: records[i][3],
			Psu: records[i][4],
			SidePanel: records[i][5],
			ExternalVolume: utils.ToFloat64(records[i][6]),
			Internal35Bays: utils.ToFloat64(records[i][7]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
