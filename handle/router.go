package handle

import (
	"net/http"

	"github.com/gorilla/context"
)

// RegisterRouter ...
func RegisterRouter() {
	http.HandleFunc("/", middleAuth(indexHandler))
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/logout", middleAuth(logoutHandler))
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
