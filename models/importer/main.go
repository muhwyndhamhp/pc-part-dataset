package main

import (

	"fmt"

	"github.com/muhwyndhamhp/pc-part-dataset/config"
	"github.com/muhwyndhamhp/pc-part-dataset/models"
	libsql "github.com/renxzen/gorm-libsql"
	"gorm.io/gorm"

	"golang.org/x/sync/errgroup"
)

func main() {
	url := config.Get(config.LIBSQL_URL)
	auth := config.Get(config.LIBSQL_TOKEN)

	db, err := gorm.Open(libsql.Open(fmt.Sprintf("%s?authToken=%s", url, auth)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to parent database: %s\n", db.Name())

	group := errgroup.Group{}

	
	group.Go(func() error {
		err = models.CaseAccessory{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.CaseFan{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Case{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.CpuCooler{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.CPU{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.ExternalHardDrive{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.FanController{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Headphones{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.InternalHardDrive{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Keyboard{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Memory{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Monitor{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Motherboard{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Mouse{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.OpticalDrive{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.OS{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.PowerSupply{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.SoundCard{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Speakers{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.ThermalPaste{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.UPS{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.VideoCard{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.Webcam{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.WiredNetworkCard{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	
	group.Go(func() error {
		err = models.WirelessNetworkCard{}.ImportData(db)
		if err != nil {
			return err
		}
		return nil
	})
	

	if err := group.Wait(); err != nil {
		panic(err)
	}
}
