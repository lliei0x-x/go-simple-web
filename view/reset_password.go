package view

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
