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
func (IndexVMInstance) GetIndexVM(username string) IndexVM {
	user, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(user.ID)

	vm := IndexVM{BaseViewModel{Title: "HomePage"}, *user, *posts}
	vm.setCurrentUser(username)
	return vm
}
