package view

import (
	"go-simple-web/model"
)

// ResetPWDVM struct
type ResetPWDVM struct {
	LoginVM
	Token string
}

// ResetPWDVMInstance struct
type ResetPWDVMInstance struct{}

// GetVM func
func (ResetPWDVMInstance) GetVM(token string) ResetPWDVM {
	vm := ResetPWDVM{}

	vm.setTitle("Reset Password")
	vm.Token = token

	return vm
}

// CheckToken func
func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

// ResetUserPassword func
func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
