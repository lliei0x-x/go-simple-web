package view

import (
	"go-simple-web/model"
)

// ProfileEditVM struct
type ProfileEditVM struct {
	BaseViewModel
	ProfileUser model.User
}

// ProfileEditVMInstance struct
type ProfileEditVMInstance struct{}

// GetVM func
func (ProfileEditVMInstance) GetVM(username string) ProfileEditVM {
	user, _ := model.GetUserByUsername(username)

	vm := ProfileEditVM{}
	vm.setTitle("Profile Edit")
	vm.setCurrentUser(username)
	vm.ProfileUser = *user

	return vm
}
