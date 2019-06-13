package handle

import (
	"net/http"

	"go-web/model/vmodel"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexVM := vmodel.IndexVMInstance{}
	vm := indexVM.GetIndexVM()
	templates["index"].Execute(w, &vm)
}
