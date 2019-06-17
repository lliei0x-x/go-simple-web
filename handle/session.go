package handle

import (
	"errors"
	"log"
	"net/http"

	"go-simple-web/config"

	"github.com/gorilla/sessions"
)

var (
	sessionName string
	store       *sessions.CookieStore
)

func init() {
	s1, s2 := config.GetSessionConfig()
	store = sessions.NewCookieStore([]byte(s1))
	sessionName = s2
}

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	log.Println("val:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
	log.Println("username:", username)
	return username, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
