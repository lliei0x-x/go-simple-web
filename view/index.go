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

// GetVM func
func (IndexVMInstance) GetVM(username string) IndexVM {
	user, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(user.ID)

	vm := IndexVM{}
	vm.setTitle("HomePage")
	vm.User = *user
	vm.Posts = *posts
	vm.setCurrentUser(username)

	return vm
}
