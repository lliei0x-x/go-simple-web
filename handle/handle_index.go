package handle

import (
	"log"
	"net/http"

	"go-simple-web/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexVM := view.IndexVMInstance{}
	username, _ := getSessionUser(r)
	flash := getFlash(w, r)
	page := getPage(r)
	vm := indexVM.GetVM(username, flash, page, pageLimit)
	if r.Method == http.MethodGet {
		templates["index"].Execute(w, &vm)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		body := r.Form.Get("body")
		errMessage := checkLen("Post", body, 1, 180)
		if errMessage != "" {
			setFlash(w, r, errMessage)
		} else {
			err := view.CreatePost(username, body)
			if err != nil {
				log.Println("add Post error:", err)
				w.Write([]byte("Error insert Post in database"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
