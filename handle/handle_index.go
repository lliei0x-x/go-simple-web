package handle

import (
	"net/http"

	"go-simple-web/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexVM := view.IndexVMInstance{}
	username, _ := getSessionUser(r)
	vm := indexVM.GetIndexVM(username)
	templates["index"].Execute(w, &vm)
}
