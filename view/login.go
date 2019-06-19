package view

// LoginVM login View Model
type LoginVM struct {
	BaseViewModel
	ErrInfos []string
}

// LoginVMInstance struct
type LoginVMInstance struct{}

// GetVM func
func (LoginVMInstance) GetVM() LoginVM {
	vm := LoginVM{}
	vm.setTitle("Login")
	return vm
}

// AddErrInfo func
func (v *LoginVM) AddErrInfo(errs ...string) {
	v.ErrInfos = append(v.ErrInfos, errs...)
}
