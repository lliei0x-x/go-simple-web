package view

import (
	"go-simple-web/config"
	"go-simple-web/model"
)

// EmailVM struct
type EmailVM struct {
	Username string
	Token    string
	Server   string
}

// EmailVMInstance struct
type EmailVMInstance struct{}

// GetVM func
func (EmailVMInstance) GetVM(email string) EmailVM {
	vm := EmailVM{}
	user, _ := model.GetUserByEmail(email)
	vm.Username = user.Username
	vm.Token, _ = user.GenerateToken()
	vm.Server = config.GetServerURL()
	return vm
}
