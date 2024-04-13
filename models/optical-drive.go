package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed optical-drive.csv
var OpticalDriveCsv []byte

type OpticalDrive struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Bd int `json:"bd" form:"bd"`
	 Dvd int `json:"dvd" form:"dvd"`
	 Cd int `json:"cd" form:"cd"`
	 BdWrite string `json:"bd_write" form:"bd_write"`
	 DvdWrite string `json:"dvd_write" form:"dvd_write"`
	 CdWrite string `json:"cd_write" form:"cd_write"`
	
}

func (OpticalDrive)TableName() string {
	return "optical_drives"
}

func (OpticalDrive)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(OpticalDriveCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*OpticalDrive{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := OpticalDrive{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Bd: utils.ToInt(records[i][2]),
			Dvd: utils.ToInt(records[i][3]),
			Cd: utils.ToInt(records[i][4]),
			BdWrite: records[i][5],
			DvdWrite: records[i][6],
			CdWrite: records[i][7],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
