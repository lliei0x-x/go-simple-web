package handle

import (
	"fmt"
	"log"
	"net/http"

	"go-simple-web/view"
	"go-simple-web/model"
)

func profileEditHandler(w http.ResponseWriter, r *http.Request) {
	username, _ := getSessionUser(r)
	profileEditVM := view.ProfileEditVMInstance{}
	vm := profileEditVM.GetVM(username)
	if r.Method == http.MethodGet {
		templates["profile_edit"].Execute(w, &vm)
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		text := r.Form.Get("aboutme")
		log.Println(text)
		err := model.UpdateAboutMe(username,text)
		if err != nil {
			log.Printf("update aboutme by username[%v] err[%v]\n", username, err)
			w.Write([]byte("Error update aboutme"))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/user/%s", username), http.StatusSeeOther)
	}
}