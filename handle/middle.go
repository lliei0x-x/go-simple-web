package handle

import (
	"log"
	"net/http"

	"go-simple-web/model"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		if username != "" {
			log.Println("Last seen:", username)
			err := model.UpdateLastSeen(username)
			if err != nil {
				log.Printf("update lastseen by username[%v] err[%v]", username, err)
			}
		}
		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
