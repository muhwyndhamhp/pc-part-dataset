package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed webcam.csv
var WebcamCsv []byte

type Webcam struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Resolutions string `json:"resolutions" form:"resolutions"`
	 Connection string `json:"connection" form:"connection"`
	 FocusType string `json:"focus_type" form:"focus_type"`
	 Os string `json:"os" form:"os"`
	 Fov float64 `json:"fov" form:"fov"`
	
}

func (Webcam)TableName() string {
	return "webcams"
}

func (Webcam)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(WebcamCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*Webcam{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := Webcam{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Resolutions: records[i][2],
			Connection: records[i][3],
			FocusType: records[i][4],
			Os: records[i][5],
			Fov: utils.ToFloat64(records[i][6]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
