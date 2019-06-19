package handle

import (
	"log"
	"net/http"

	"go-simple-web/model"
	"go-simple-web/view"
)

func registerHandle(w http.ResponseWriter, r *http.Request) {
	registerVM := view.RegisterVMInstance{}
	vm := registerVM.GetVM()
	if r.Method == http.MethodGet {
		templates["register"].Execute(w, &vm)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd := r.Form.Get("pwd")
		repwd := r.Form.Get("repwd")

		errs := checkRegister(username, email, pwd, repwd)
		log.Println(errs)
		if len(errs) > 0 {
			vm.AddErrInfo(errs...)
			templates["register"].Execute(w, &vm)
		} else {
			if err := model.AddUser(username, pwd, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}
