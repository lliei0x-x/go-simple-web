package view

import "go-simple-web/model"

// ProfileVM struct
type ProfileVM struct {
	BaseViewModel
	Posts       []model.Post
	ProfileUser model.User
}

// ProfileVMInstance struct
type ProfileVMInstance struct{}

// GetProfileVM func
func (ProfileVMInstance) GetProfileVM(cUser, pUser string) ProfileVM {
	user, _ := model.GetUserByUsername(pUser)
	posts, _ := model.GetPostsByUserID(user.ID)

	vm := ProfileVM{}
	vm.ProfileUser = *user
	vm.Posts = *posts
	vm.setTitle("Profile")
	vm.setCurrentUser(cUser)

	return vm
}
