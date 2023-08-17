package admincontroller

import (
	"log"
	"net/http"
	"text/template"
)

func AdminController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/index.html",
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

func AdminLoginController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-signin.html",
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

func AdminProfileController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-profile.html",
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

func AdminChangePasswordController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-change-password.html",
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

func AdminForgotPasswordController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-forgot-password.html",
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
