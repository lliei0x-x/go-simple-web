package handle

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"go-simple-web/model"
)

// Login Check
func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

func checkPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}

func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}

func checkUserPassword(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Printf("Can not find username[%v] err[%v]\n", username, err)
		return false
	}
	return user.CheckPassword(password)
}

func checkUserExist(username string) bool {
	user, _ := model.GetUserByUsername(username)
	if user != nil {
		return true
	}
	return false
}

func checkLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserPassword(username, password); !errCheck {
		errs = append(errs, "Username or password is not correct")
	}
	return errs
}

func checkRegister(username, email, pwd, repwd string) []string {
	var errs []string
	if pwd != repwd {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(pwd); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); errCheck {
		errs = append(errs, "Username already exist, please choose another username")
	}
	return errs
}

func getPage(r *http.Request) int {
	url := r.URL         // net/url.URL
	query := url.Query() // Values (map[string][]string)

	q := query.Get("page")
	if q == "" {
		return 1
	}

	page, err := strconv.Atoi(q)
	if err != nil {
		return 1
	}
	return page
}
