package handle

import (
	"net/http"

	"go-simple-web/view"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginVM := view.LoginVMInstance{}
	vm := loginVM.GetVM()
	if r.Method == http.MethodGet {
		templates["login"].Execute(w, &vm)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		errs := checkLogin(username, password)
		if len(errs) > 0 {
			vm.AddErrInfo(errs...)
			templates["login"].Execute(w, &vm)
		} else {
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
