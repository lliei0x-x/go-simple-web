package handle

import (
	"net/http"
)

func followHandler(w http.ResponseWriter, r *http.Request) {
	usrnamegetSessionUser(r)
}
func unFollowHandler(w http.ResponseWriter, r *http.Request) {}
