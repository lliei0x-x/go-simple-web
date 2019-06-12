package handle

import (
	"net/http"

	"go-web/model"
)

var errInfos []string

func loginHandle(w http.ResponseWriter, r *http.Request) {
	loginVM := model.LoginVMInstance{}
	vm := loginVM.GetLoginVM()
	if r.Method == http.MethodGet {
		templates["login"].Execute(w, &vm)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			addErrInfo("username must longer than 3")
		}

		if len(password) < 6 {
			addErrInfo("password must longer than 6")
		}

		if !check(username, password) {
			addErrInfo("username password not correct, please input again")
		}

		if len(errInfos) > 0 {
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

// AddErrInfo ...
func addErrInfo(errInfo string) {
	errInfos = append(errInfos, errInfo)
}
