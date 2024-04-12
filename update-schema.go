package main

import (
	"fmt"
	"go/importer"

	"github.com/muhwyndhamhp/pc-part-dataset/models"
)

// "github.com/muhwyndhamhp/pc-part-dataset/config"
// libsql "github.com/renxzen/gorm-libsql"
// "gorm.io/gorm"

func UpdateParentDBSchema() {
	// url := config.Get(config.LIBSQL_URL)
	// auth := config.Get(config.LIBSQL_TOKEN)
	//
	// db, err := gorm.Open(libsql.Open(fmt.Sprintf("%s?authToken=%s", url, auth)), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Printf("Connected to parent database: %s\n", db.Name())

	pkg, err := importer.Default().Import("github.com/muhwyndhamhp/pc-part-dataset/models")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	_ = models.Base{}

	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
}
