package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed os.csv
var OSCsv []byte

type OS struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Mode int `json:"mode" form:"mode"`
	 MaxMemory int `json:"max_memory" form:"max_memory"`
	
}

func (OS)TableName() string {
	return "os"
}

func (OS)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(OSCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*OS{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := OS{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Mode: utils.ToInt(records[i][2]),
			MaxMemory: utils.ToInt(records[i][3]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
