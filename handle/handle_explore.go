package handle

import (
	"go-simple-web/view"
	"net/http"
)

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	exploreVM := view.ExploreVMInstance{}
	username, _ := getSessionUser(r)
	page := getPage(r)
	vm := exploreVM.GetVM(username, page, pageLimit)
	templates["explore"].Execute(w, &vm)
}
