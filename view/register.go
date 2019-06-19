package view

// RegisterVM struct
type RegisterVM struct {
	BaseViewModel
	ErrInfos []string
}

// RegisterVMInstance struct
type RegisterVMInstance struct{}

// GetVM func
func (RegisterVMInstance) GetVM() RegisterVM {
	vm := RegisterVM{}
	vm.setTitle("Register")
	return vm
}

// AddErrInfo func
func (v *RegisterVM) AddErrInfo(errs ...string) {
	v.ErrInfos = append(v.ErrInfos, errs...)
}
