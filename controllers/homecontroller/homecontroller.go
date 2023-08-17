package homecontroller

import (
	"log"
	"net/http"
	"text/template"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/client/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
