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
	page := getPage(r)

	profileVM := view.ProfileVMInstance{}
	cUser, _ := getSessionUser(r)
	vm := profileVM.GetVM(cUser, pUser, page, pageLimit)

	templates["profile"].Execute(w, &vm)
}
