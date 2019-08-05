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
		layout = template.Must(template.ParseFiles(dirPath + "/_base.html"))
		if !strings.Contains(v, "_base.html") {
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

//if !strings.Contains(v, "404.html")
