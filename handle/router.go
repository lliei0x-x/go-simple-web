package handle

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// RegisterRouter ...
func RegisterRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleAuth(indexHandler))
	r.HandleFunc("/{username}", middleAuth(ProfileHandler))
	r.HandleFunc("/login", loginHandle)
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	http.Handle("/", r)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))

}
