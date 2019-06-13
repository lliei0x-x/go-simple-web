package main

import (
	"go-web/config"
	"go-web/handle"
	"go-web/model"
)

func main() {

	// create table by gorm auto migrate
	createTable()
	defer config.POSTGRES.Close()

	handle.RegisterRouter()

}

func createTable() {
	config.POSTGRES.AutoMigrate(
		&model.User{},
		&model.Post{},
	)
}
