package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed monitor.csv
var MonitorCsv []byte

type Monitor struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 ScreenSize int `json:"screen_size" form:"screen_size"`
	 Resolution string `json:"resolution" form:"resolution"`
	 RefreshRate int `json:"refresh_rate" form:"refresh_rate"`
	 ResponseTime float64 `json:"response_time" form:"response_time"`
	 PanelType string `json:"panel_type" form:"panel_type"`
	 AspectRatio string `json:"aspect_ratio" form:"aspect_ratio"`
	
}

func (Monitor)TableName() string {
	return "monitors"
}

func (Monitor)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(MonitorCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Monitor{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Monitor{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			ScreenSize: utils.ToInt(records[i][2]),
			Resolution: records[i][3],
			RefreshRate: utils.ToInt(records[i][4]),
			ResponseTime: utils.ToFloat64(records[i][5]),
			PanelType: records[i][6],
			AspectRatio: records[i][7],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
