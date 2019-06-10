package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"go-web/common"
	"go-web/model"
)

const (
	dirPath = "./template"
)

func main() {
	var (
		layout, tmpl *template.Template
	)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// tpl, _ := template.ParseFiles("./template/conent/index.html")

		var files []string
		files, err := common.GetAllFiles(dirPath, files, ".html")
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range files {
			if strings.Contains(v, "_base.html") {
				layout = template.Must(template.ParseFiles(v))
			} else {

				indexVM := model.IndexVMInstance{}
				ivm := indexVM.GetIndexVM()

				tmpl = template.Must(layout.Clone())
				_, err = tmpl.ParseFiles(v)
				tmpl.Execute(w, &ivm)
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
