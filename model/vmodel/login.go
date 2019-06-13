package vmodel

import (
	"go-web/model/basic"
)

// LoginVM login View Model
type LoginVM struct {
	basic.BasicTitle
	ErrInfos []string
}

// LoginVMInstance ...
type LoginVMInstance struct{}

// GetLoginVM ...
func (LoginVMInstance) GetLoginVM() LoginVM {
	vm := LoginVM{}
	vm.SetBasicTitle("Login")
	return vm
}

// AddErrInfo ...
func (v *LoginVM) AddErrInfo(s string) {
	v.ErrInfos = append(v.ErrInfos, s)
}
