package handle

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"go-simple-web/view"
)

func resetPWDHandle(w http.ResponseWriter, r *http.Request) {
	resetPWDVM := view.RegisterVMInstance{}
	vm := resetPWDVM.GetVM()

	if r.Method == http.MethodGet {
		templates["reset_password"].Execute(w, &vm)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.Form.Get("email")

		errs := checkResetPassword(email)
		vm.AddErrInfo(errs...)

		if len(errs) > 0 {
			templates["reset_password"].Execute(w, &vm)

		} else {
			log.Println("Send mail to", email)
			emailVM := view.EmailVMInstance{}
			vm := emailVM.GetVM(email)
			var contentByte bytes.Buffer
			tpl, _ := template.ParseFiles("templates/email.html")

			if err := tpl.Execute(&contentByte, &vm); err != nil {
				log.Println("Get Parse Template:", err)
				w.Write([]byte("Error send email"))
				return
			}
			content := contentByte.String()
			go sendEmail(email, "Reset Password", content)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
