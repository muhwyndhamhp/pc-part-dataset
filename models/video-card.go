package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed video-card.csv
var VideoCardCsv []byte

type VideoCard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Chipset string `json:"chipset" form:"chipset"`
	 Memory float64 `json:"memory" form:"memory"`
	 CoreClock float64 `json:"core_clock" form:"core_clock"`
	 BoostClock float64 `json:"boost_clock" form:"boost_clock"`
	 Color string `json:"color" form:"color"`
	 Length float64 `json:"length" form:"length"`
	
}

func (VideoCard)TableName() string {
	return "video_cards"
}

func (VideoCard)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(VideoCardCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*VideoCard{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := VideoCard{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Chipset: records[i][2],
			Memory: utils.ToFloat64(records[i][3]),
			CoreClock: utils.ToFloat64(records[i][4]),
			BoostClock: utils.ToFloat64(records[i][5]),
			Color: records[i][6],
			Length: utils.ToFloat64(records[i][7]),
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
