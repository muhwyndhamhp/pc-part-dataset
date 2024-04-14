package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed external-hard-drive.csv
var ExternalHardDriveCsv []byte

type ExternalHardDrive struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Type string `json:"type" form:"type"`
	 Interface string `json:"interface" form:"interface"`
	 Capacity float64 `json:"capacity" form:"capacity"`
	 PricePerGb float64 `json:"price_per_gb" form:"price_per_gb"`
	 Color string `json:"color" form:"color"`
	
}

func (ExternalHardDrive)TableName() string {
	return "external_hard_drives"
}

func (ExternalHardDrive)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(ExternalHardDriveCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*ExternalHardDrive{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := ExternalHardDrive{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Type: records[i][2],
			Interface: records[i][3],
			Capacity: utils.ToFloat64(records[i][4]),
			PricePerGb: utils.ToFloat64(records[i][5]),
			Color: records[i][6],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
