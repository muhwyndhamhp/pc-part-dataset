package main

import (
	"fmt"

	"github.com/muhwyndhamhp/pc-part-dataset/config"
	"github.com/muhwyndhamhp/pc-part-dataset/models"
	libsql "github.com/renxzen/gorm-libsql"
	"gorm.io/gorm"
)

func main() {
	url := config.Get(config.LIBSQL_SCHEMA_URL)
	auth := config.Get(config.LIBSQL_SCHEMA_TOKEN)

	db, err := gorm.Open(libsql.Open(fmt.Sprintf("%s?authToken=%s", url, auth)), &gorm.Config{
		DryRun: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to parent database: %s\n", db.Name())

	
		err = db.AutoMigrate(&models.CaseAccessory{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.CaseFan{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Case{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.CpuCooler{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.CPU{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.ExternalHardDrive{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.FanController{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Headphones{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.InternalHardDrive{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Keyboard{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Memory{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Monitor{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Motherboard{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Mouse{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.OpticalDrive{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.OS{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.PowerSupply{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.SoundCard{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Speakers{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.ThermalPaste{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.UPS{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.VideoCard{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.Webcam{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.WiredNetworkCard{})
		if err != nil {panic(err)}
	
		err = db.AutoMigrate(&models.WirelessNetworkCard{})
		if err != nil {panic(err)}
	
}
