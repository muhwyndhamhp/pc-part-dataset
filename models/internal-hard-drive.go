package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed internal-hard-drive.csv
var InternalHardDriveCsv []byte

type InternalHardDrive struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Capacity int `json:"capacity" form:"capacity"`
	 PricePerGb float64 `json:"price_per_gb" form:"price_per_gb"`
	 Type string `json:"type" form:"type"`
	 Cache int `json:"cache" form:"cache"`
	 FormFactor string `json:"form_factor" form:"form_factor"`
	 Interface string `json:"interface" form:"interface"`
	
}

func (InternalHardDrive)TableName() string {
	return "internal_hard_drives"
}

func (InternalHardDrive)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(InternalHardDriveCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*InternalHardDrive{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := InternalHardDrive{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Capacity: utils.ToInt(records[i][2]),
			PricePerGb: utils.ToFloat64(records[i][3]),
			Type: records[i][4],
			Cache: utils.ToInt(records[i][5]),
			FormFactor: records[i][6],
			Interface: records[i][7],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
