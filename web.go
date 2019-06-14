package main

import (
	"go-simple-web/common"
	"go-simple-web/config"
	"go-simple-web/handle"
	"go-simple-web/model"
)

func main() {

	// create table by gorm auto migrate
	createTable()
	defer config.POSTGRES.Close()
	testData()

	// register
	handle.RegisterRouter()

}

func createTable() {
	config.POSTGRES.AutoMigrate(
		&model.User{},
		&model.Post{},
	)
}

func testData() {
	users := []model.User{
		{
			Username:     "bonfy",
			PasswordHash: common.GeneratePasswordHash("abc123"),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
			},
		},
		{
			Username:     "rene",
			PasswordHash: common.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Posts: []model.Post{
				{Body: "The Avengers movie was so cool!"},
				{Body: "Sun shine is beautiful"},
			},
		},
	}
	for _, u := range users {
		config.POSTGRES.Debug().Create(&u)
	}
}
