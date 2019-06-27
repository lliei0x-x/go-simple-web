package view

import (
	"go-simple-web/model"
)

// IndexVM Index View Model
type IndexVM struct {
	BaseViewModel
	model.User
	Posts []model.Post
	Flash string
}

// IndexVMInstance ...
type IndexVMInstance struct{}

// GetVM func
func (IndexVMInstance) GetVM(username, flash string) IndexVM {
	user, _ := model.GetUserByUsername(username)
	// posts, _ := model.GetPostsByUserID(user.ID)
	posts, _ := user.FollowingPosts()
	vm := IndexVM{}
	vm.setTitle("HomePage")
	vm.User = *user
	vm.Posts = *posts
	vm.setCurrentUser(username)
	vm.Flash = flash

	return vm
}

// CreatePost func
func CreatePost(username, body string) error {
	user, _ := model.GetUserByUsername(username)
	return user.CreatePost(body)
}
