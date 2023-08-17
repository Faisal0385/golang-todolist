package admintodocontroller

import (
	"log"
	"net/http"
	"text/template"
)

func AdminAddTodoController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-add-todo.html",
		"./views/admin/master_admin_layout.html",
		"./views/admin/components/footer_partial.html",
		"./views/admin/components/sidebar_partial.html",
		"./views/admin/components/navbar_partial.html",
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

func AdminListTodoController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-todo-list.html",
		"./views/admin/master_admin_layout.html",
		"./views/admin/components/footer_partial.html",
		"./views/admin/components/sidebar_partial.html",
		"./views/admin/components/navbar_partial.html",
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
