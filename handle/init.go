package handle

import (
	"html/template"
	"log"
	"strings"

	"go-simple-web/common"
)

var templates map[string]*template.Template

const (
	dirPath   = "./template"
	pageLimit = 5
)

func init() {
	getTemplates()
}

func getTemplates() {
	var (
		layout, tmpl *template.Template
		files        []string
		err          error
	)

	if files, err = common.GetAllFiles(dirPath, files, ".html"); err != nil {
		log.Fatal(err)
	}

	templatesMap := make(map[string]*template.Template)
	for _, v := range files {
		if strings.Contains(v, "_base.html") {
			layout = template.Must(template.ParseFiles(v))
		} else {
			tmpl = template.Must(layout.Clone())
			if _, err = tmpl.ParseFiles(v); err != nil {
				log.Fatal(err)
			}
			n := common.MatchModelName(v)
			templatesMap[n] = tmpl
		}
	}
	templates = templatesMap
}
