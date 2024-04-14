package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed mouse.csv
var MouseCsv []byte

type Mouse struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 TrackingMethod string `json:"tracking_method" form:"tracking_method"`
	 ConnectionType string `json:"connection_type" form:"connection_type"`
	 MaxDpi float64 `json:"max_dpi" form:"max_dpi"`
	 HandOrientation string `json:"hand_orientation" form:"hand_orientation"`
	 Color string `json:"color" form:"color"`
	
}

func (Mouse)TableName() string {
	return "mouses"
}

func (Mouse)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(MouseCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Mouse{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Mouse{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			TrackingMethod: records[i][2],
			ConnectionType: records[i][3],
			MaxDpi: utils.ToFloat64(records[i][4]),
			HandOrientation: records[i][5],
			Color: records[i][6],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
