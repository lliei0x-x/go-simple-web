package view

import "go-simple-web/model"

// ExploreViewModel struct
type ExploreViewModel struct {
	BaseViewModel
	PageViewModel
	Posts []model.Post
}

// ExploreVMInstance struct
type ExploreVMInstance struct{}

// GetVM func
func (ExploreVMInstance) GetVM(username string, page, limit int) ExploreViewModel {

	posts, total, _ := model.GetPostsByPageAndLimit(page, limit)

	vm := ExploreViewModel{}
	vm.Posts = *posts
	vm.setTitle("Explore")
	vm.setCurrentUser(username)
	vm.setPageViewModel(total, page, limit)

	return vm
}
