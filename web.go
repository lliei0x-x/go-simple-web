package main

import (
	"go-simple-web/config"
	"go-simple-web/handle"
	"go-simple-web/model"
)

func main() {

	// create table by gorm auto migrate
	createTable()
	defer config.POSTGRES.Close()

	// register
	handle.RegisterRouter()

}

func createTable() {
	config.POSTGRES.AutoMigrate(
		&model.User{},
		&model.Post{},
	)
}
