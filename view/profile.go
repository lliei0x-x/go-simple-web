package view

import (
	"go-simple-web/model"
	"time"
)

// ProfileVM struct
type ProfileVM struct {
	BaseViewModel
	Editable    bool
	Posts       []model.Post
	ProfileUser model.User
}

// ProfileVMInstance struct
type ProfileVMInstance struct{}

// GetVM func
func (ProfileVMInstance) GetVM(cUser, pUser string) ProfileVM {
	user, _ := model.GetUserByUsername(pUser)
	posts, _ := model.GetPostsByUserID(user.ID)

	vm := &ProfileVM{}
	vm.setTitle("Profile")
	vm.Editable = (cUser == pUser)
	vm.Posts = *posts
	vm.ProfileUser = *user
	vm.setTime()
	vm.setCurrentUser(cUser)

	return *vm
}

func (v *ProfileVM) setTime() {
	v.ProfileUser.LastSeen = v.ProfileUser.LastSeen.In(time.FixedZone("CST", 8*3600))
}
