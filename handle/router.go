package handle

import "net/http"

// RegisterRouter ...
func RegisterRouter() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandle)
	http.ListenAndServe(":8080", nil)
}
