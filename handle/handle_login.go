package handle

import (
	"log"
	"net/http"

	"go-simple-web/model"
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

		if len(username) < 3 {
			vm.AddErrInfo("username must longer than 3")
		}

		if len(password) < 6 {
			vm.AddErrInfo("password must longer than 6")
		}

		if !checkLogin(username, password) {
			vm.AddErrInfo("username password not correct, please input again")
		}

		if len(vm.ErrInfos) > 0 {
			templates["login"].Execute(w, &vm)
		} else {
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func checkLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password)
}
