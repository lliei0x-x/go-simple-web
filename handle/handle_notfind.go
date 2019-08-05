package handle

import (
	"log"
	"net/http"

	"go-simple-web/view"
)

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	flash := getFlash(w, r)
	log.Println("not find")
	vm := view.NotFoundMessage{
		Flash: flash,
	}
	templates["404"].Execute(w, &vm)
}
