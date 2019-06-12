package model

import (
	"go-web/model/basic"
)

// LoginVM login View Model
type LoginVM struct {
	basic.BasicTitle
}

// LoginVMInstance ...
type LoginVMInstance struct{}

// GetLoginVM ...
func (LoginVMInstance) GetLoginVM() LoginVM {
	vm := LoginVM{}
	vm.SetBasicTitle("Login")
	return vm
}
