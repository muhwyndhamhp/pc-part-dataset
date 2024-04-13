package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed wired-network-card.csv
var WiredNetworkCardCsv []byte

type WiredNetworkCard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Interface string `json:"interface" form:"interface"`
	 Color string `json:"color" form:"color"`
	
}

func (WiredNetworkCard)TableName() string {
	return "wired_network_cards"
}

func (WiredNetworkCard)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(WiredNetworkCardCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*WiredNetworkCard{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := WiredNetworkCard{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Interface: records[i][2],
			Color: records[i][3],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
