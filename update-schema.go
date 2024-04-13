package main

import (
	"fmt"

	"github.com/muhwyndhamhp/pc-part-dataset/config"
	libsql "github.com/renxzen/gorm-libsql"
	"gorm.io/gorm"
)

func UpdateParentDBSchema() {
	url := config.Get(config.LIBSQL_SCHEMA_URL)
	auth := config.Get(config.LIBSQL_SCHEMA_TOKEN)

	db, err := gorm.Open(libsql.Open(fmt.Sprintf("%s?authToken=%s", url, auth)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to parent database: %s\n", db.Name())
}
