package handle

import (
	"log"
	"net/http"

	"go-simple-web/view"

	"github.com/gorilla/mux"
)

func resetPWDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	username, err := view.CheckToken(token)
	if err != nil {
		w.Write([]byte("The token is no longer valid, please go to the login page."))
	}

	resetPWDVM := view.ResetPWDVMInstance{}
	vm := resetPWDVM.GetVM(token)

	if r.Method == http.MethodGet {
		templates["reset_pwd"].Execute(w, &vm)
	}

	if r.Method == http.MethodPost {
		log.Println("Reset password for ", username)
		r.ParseForm()
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkResetPassword(pwd1, pwd2)
		vm.AddErrInfo(errs...)

		if len(vm.ErrInfos) > 0 {
			templates["reset_pwd"].Execute(w, &vm)
		} else {
			if err := view.ResetUserPassword(username, pwd1); err != nil {
				log.Println("reset User password error:", err)
				w.Write([]byte("Error update user password in database"))
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
