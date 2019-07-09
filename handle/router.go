package handle

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// RegisterRouter ...
func RegisterRouter() {
	r := mux.NewRouter() // 匹配

	r.HandleFunc("/", middleAuth(indexHandler))
	r.HandleFunc("/explore", middleAuth(exploreHandler))
	r.HandleFunc("/register", registerHandle)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	r.HandleFunc("/user/{username}", middleAuth(profileHandler))
	r.HandleFunc("/profile_edit", middleAuth(profileEditHandler))
	r.HandleFunc("/follow/{username}", middleAuth(followHandler))
	r.HandleFunc("/unfollow/{username}", middleAuth(unFollowHandler))
	http.Handle("/", r)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))

}
