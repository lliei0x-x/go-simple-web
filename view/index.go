package view

import (
	"go-simple-web/model"
)

// IndexVM Index View Model
type IndexVM struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexVMInstance ...
type IndexVMInstance struct{}

// GetIndexVM ...
func (IndexVMInstance) GetIndexVM() IndexVM {
	user, _ := model.GetUserByUsername("rene")
	posts, _ := model.GetPostsByUserID(user.ID)

	vm := IndexVM{BaseViewModel{Title: "HomePage"}, *user, *posts}
	return vm
}
