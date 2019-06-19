package main

import (
	"fmt"

	"go-simple-web/common"
	"go-simple-web/config"
	"go-simple-web/handle"
	"go-simple-web/model"
)

func main() {

	// create table by gorm auto migrate
	createTable()
	defer config.POSTGRES.Close()

	// testDBData()

	// register
	handle.RegisterRouter()

}

func createTable() {
	config.POSTGRES.AutoMigrate(
		&model.User{},
		&model.Post{},
	)
}

func testDBData() {
	users := []model.User{
		{
			Username:     "leeifme",
			PasswordHash: common.GeneratePasswordHash("leeifme"),
			Email:        "lee@test.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=retro", common.Md5("lee@test.com")),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
			},
		},
		{
			Username:     "rene",
			PasswordHash: common.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", common.Md5("rene@test.com")),
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
