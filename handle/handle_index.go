package handle

import (
	"net/http"

	"go-web/model"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexVM := model.IndexVMInstance{}
	vm := indexVM.GetIndexVM()
	templates["index"].Execute(w, &vm)
}
