package handle

import (
	"net/http"

	"go-simple-web/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexVM := view.IndexVMInstance{}
	vm := indexVM.GetIndexVM()
	templates["index"].Execute(w, &vm)
}
