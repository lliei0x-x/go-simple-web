package handle

import (
	"fmt"
	"log"
	"net/http"

	"go-simple-web/model"

	"github.com/gorilla/mux"
)

func followHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	other := vars["username"]
	username, _ := getSessionUser(r)
	cUser, _ := model.GetUserByUsername(username)
	if err := cUser.Follow(other); err != nil {
		log.Println("Follow error:", err)
		w.Write([]byte("Error in Follow"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%v", other), http.StatusSeeOther)
}

func unFollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	other := vars["username"]
	username, _ := getSessionUser(r)
	cUser, _ := model.GetUserByUsername(username)
	if err := cUser.UnFollow(other); err != nil {
		log.Println("UnFollow error:", err)
		w.Write([]byte("Error in UnFollow"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%v", other), http.StatusSeeOther)
}
