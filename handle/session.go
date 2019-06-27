package handle

import (
	"errors"
	"fmt"
	"net/http"

	"go-simple-web/config"

	"github.com/gorilla/sessions"
)

var (
	sessionName string
	flashName   string
	store       *sessions.CookieStore
)

func init() {
	s1, s2, s3 := config.GetSessionConfig()
	store = sessions.NewCookieStore([]byte(s1))
	sessionName = s2
	flashName = s3
}

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
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

func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := store.Get(r, sessionName)
	session.AddFlash(message, flashName)
	session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}
	session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}
