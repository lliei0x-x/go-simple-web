package vmodel

import (
	"go-web/model/basic"
)

// IndexVM Index View Model
type IndexVM struct {
	basic.BasicTitle
	basic.User
	Posts []basic.Post
}

// IndexVMInstance ...
type IndexVMInstance struct{}

// GetIndexVM ...
func (IndexVMInstance) GetIndexVM() IndexVM {
	u1 := basic.User{Username: "bonfy"}
	u2 := basic.User{Username: "rene"}

	posts := []basic.Post{
		basic.Post{User: u1, Body: "Beautiful day in Portland!"},
		basic.Post{User: u2, Body: "The Avengers movie was so cool!"},
	}

	vm := IndexVM{basic.BasicTitle{Title: "leeifme"}, u1, posts}
	return vm
}
