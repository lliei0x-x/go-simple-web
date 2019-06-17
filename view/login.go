package view

// LoginVM login View Model
type LoginVM struct {
	BaseViewModel
	ErrInfos []string
}

// LoginVMInstance ...
type LoginVMInstance struct{}

// GetLoginVM ...
func (LoginVMInstance) GetLoginVM() LoginVM {
	vm := LoginVM{}
	vm.setTitle("Login")
	return vm
}

// AddErrInfo ...
func (v *LoginVM) AddErrInfo(s string) {
	v.ErrInfos = append(v.ErrInfos, s)
}
