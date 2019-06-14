package handle

import (
	"net/http"

	"go-simple-web/view"
)

func loginHandle(w http.ResponseWriter, r *http.Request) {
	loginVM := view.LoginVMInstance{}
	vm := loginVM.GetLoginVM()
	if r.Method == http.MethodGet {
		templates["login"].Execute(w, &vm)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			vm.AddErrInfo("username must longer than 3")
		}

		if len(password) < 6 {
			vm.AddErrInfo("password must longer than 6")
		}

		if !check(username, password) {
			vm.AddErrInfo("username password not correct, please input again")
		}

		if len(vm.ErrInfos) > 0 {
			templates["login"].Execute(w, &vm)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func check(username, password string) bool {
	if username == "leeifme" && password == "123456" {
		return true
	}
	return false
}
