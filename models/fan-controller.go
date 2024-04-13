package models

import (
	_ "embed"

	"bytes"
	"context"
	"gorm.io/gorm"
	"encoding/csv"
	"github.com/muhwyndhamhp/pc-part-dataset/utils"
)

//go:embed fan-controller.csv
var FanControllerCsv []byte

type FanController struct {
	gorm.Model
	 Name string `json:"name" form:"name"`
	 Price float64 `json:"price" form:"price"`
	 Channels int `json:"channels" form:"channels"`
	 ChannelWattage string `json:"channel_wattage" form:"channel_wattage"`
	 Pwm bool `json:"pwm" form:"pwm"`
	 FormFactor string `json:"form_factor" form:"form_factor"`
	 Color string `json:"color" form:"color"`
	
}

func (FanController)TableName() string {
	return "fan_controllers"
}

func (FanController)ImportData(db *gorm.DB) error {
	b := bytes.NewBuffer(FanControllerCsv)
	r := csv.NewReader(b)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()
	rec := []*FanController{}
	for i := range records {
		if i == 0 {
				continue
		}

		m := FanController{
			Name: records[i][0],
			Price: utils.ToFloat64(records[i][1]),
			Channels: utils.ToInt(records[i][2]),
			ChannelWattage: records[i][3],
			Pwm: utils.ToBool(records[i][4]),
			FormFactor: records[i][5],
			Color: records[i][6],
			
		}

		rec = append(rec, &m)
	}

	if err := db.WithContext(ctx).CreateInBatches(rec, 1000).Error; err != nil {
		return err
	}

	return nil
}
