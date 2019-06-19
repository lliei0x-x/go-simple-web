package handle

import (
	"net/http"

	"go-simple-web/view"

	"github.com/gorilla/mux"
)

// ProfileHandler func
func profileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	profileVM := view.ProfileVMInstance{}
	cUser, _ := getSessionUser(r)
	vm := profileVM.GetVM(cUser, pUser)
	templates["profile"].Execute(w, &vm)
}