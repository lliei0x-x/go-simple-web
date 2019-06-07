package vmodel

import (
	"go-web/vmodel/model"
)

// IndexVM Index View Model
type IndexVM struct {
	model.BasicTitle
	model.User
	Posts []model.Post
}

// IndexVMInstance ...
type IndexVMInstance struct{}

// GetIndexVM ...
func (i *IndexVMInstance) GetIndexVM() IndexVM {
	u1 := model.User{Username: "bonfy"}
	u2 := model.User{Username: "rene"}

	posts := []model.Post{
		model.Post{User: u1, Body: "Beautiful day in Portland!"},
		model.Post{User: u2, Body: "The Avengers movie was so cool!"},
	}

	indexVM := IndexVM{model.BasicTitle{Title: "leeifme"}, u1, posts}
	return indexVM
}
