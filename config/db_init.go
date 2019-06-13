package config

import (
	"log"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// POSTGRES ...
var POSTGRES *gorm.DB

func init() {
	//db init
	connectString := getPostgreConfg()
	log.Println(connectString)
	connect, err := gorm.Open(
		"postgres",
		connectString,
	)
	connect.LogMode(true)
	if err != nil {
		log.Fatalf("connect postgres failed: %v", err)
	}
	log.Println("Login postgres database success!")
	POSTGRES = connect
}
