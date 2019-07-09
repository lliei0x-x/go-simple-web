package view

import (
	"go-simple-web/model"
)

// IndexVM Index View Model
type IndexVM struct {
	BaseViewModel
	PageViewModel
	model.User
	Posts []model.Post
	Flash string
}

// IndexVMInstance ...
type IndexVMInstance struct{}

// GetVM func
func (IndexVMInstance) GetVM(username, flash string, page, limit int) IndexVM {
	user, _ := model.GetUserByUsername(username)
	// posts, _ := model.GetPostsByUserID(user.ID)
	// posts, _ := user.FollowingPosts()
	posts, total, _ := user.FollowingPostsByPageAndLimit(page, limit)
	vm := IndexVM{}
	vm.setTitle("HomePage")
	vm.User = *user
	vm.Posts = *posts
	vm.setCurrentUser(username)
	vm.Flash = flash
	vm.setPageViewModel(total, page, limit)

	return vm
}

// CreatePost func
func CreatePost(username, body string) error {
	user, _ := model.GetUserByUsername(username)
	return user.CreatePost(body)
}
