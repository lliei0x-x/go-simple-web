package main

import (
	"net/http"

	"go-web/handle"
)

func main() {
	handle.RegisterRouter()
	http.ListenAndServe(":8080", nil)
}
