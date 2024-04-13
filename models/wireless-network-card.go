package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed wireless-network-card.csv
var WirelessNetworkCardCsv []byte

type WirelessNetworkCard struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Protocol string `json:"protocol" form:"protocol"`
	 Interface string `json:"interface" form:"interface"`
	 Color string `json:"color" form:"color"`
	
}

func (WirelessNetworkCard)TableName() string {
	return "wireless_network_cards"
}

func (WirelessNetworkCard)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(WirelessNetworkCardCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*WirelessNetworkCard{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := WirelessNetworkCard{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Protocol: records[i][2],
			Interface: records[i][3],
			Color: records[i][4],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
