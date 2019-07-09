package view

import (
	"go-simple-web/model"
	"time"
)

// ProfileVM struct
type ProfileVM struct {
	BaseViewModel
	PageViewModel
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	Posts          []model.Post
	ProfileUser    model.User
}

// ProfileVMInstance struct
type ProfileVMInstance struct{}

// GetVM func
func (ProfileVMInstance) GetVM(cUser, pUser string, page, limit int) ProfileVM {
	user, _ := model.GetUserByUsername(pUser)
	// posts, _ := model.GetPostsByUserID(user.ID)
	posts, total, _ := model.GetPostsByUserIDPageAndLimit(user.ID, page, limit)

	vm := &ProfileVM{}
	vm.setTitle("Profile")
	vm.Editable = (cUser == pUser)
	vm.Posts = *posts
	vm.ProfileUser = *user
	vm.setTime()
	vm.setCurrentUser(cUser)
	if !vm.Editable {
		vm.IsFollow = user.IsFollowedByUser(cUser)
	}
	vm.FollowersCount = user.FollowersCount()
	vm.FollowingCount = user.FollowingCount()
	vm.setPageViewModel(total, page, limit)

	return *vm
}

func (v *ProfileVM) setTime() {
	v.ProfileUser.LastSeen = v.ProfileUser.LastSeen.In(time.FixedZone("CST", 8*3600))
}
