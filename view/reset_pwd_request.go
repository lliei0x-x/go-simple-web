package view

import (
	"log"

	"go-simple-web/model"
)

// ResetPWDRequestVM struct
type ResetPWDRequestVM struct {
	LoginVM
}

// ResetPWDRequestVMInstance struct
type ResetPWDRequestVMInstance struct{}

// GetVM func
func (ResetPWDRequestVMInstance) GetVM() ResetPWDRequestVM {
	vm := ResetPWDRequestVM{}

	vm.setTitle("Forget Password")

	return vm
}

// CheckEmailExist func
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("Can not find email:", email)
		return false
	}
	return true
}
