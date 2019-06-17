package view

// BaseViewModel struct
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

// SetTitle func
func (v *BaseViewModel) setTitle(title string) {
	v.Title = title
}

// SetCurrentUser func
func (v *BaseViewModel) setCurrentUser(username string) {
	v.CurrentUser = username
}
