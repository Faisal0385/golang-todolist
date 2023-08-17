package routes

import (
	"net/http"
	"todolist/controllers/admincontroller"
	"todolist/controllers/adminteamcontroller"
	"todolist/controllers/admintodocontroller"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static"))))

	// routes
	r.HandleFunc("/", admincontroller.AdminController)

	////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////
	// admin routes
	r.HandleFunc("/admin", admincontroller.AdminController)
	r.HandleFunc("/admin-login", admincontroller.AdminLoginController)
	r.HandleFunc("/admin-profile", admincontroller.AdminProfileController)
	r.HandleFunc("/admin-change-password", admincontroller.AdminChangePasswordController)
	r.HandleFunc("/admin-forget-password", admincontroller.AdminForgotPasswordController)

	r.HandleFunc("/todo-add", admintodocontroller.AdminAddTodoController)
	r.HandleFunc("/todo-list", admintodocontroller.AdminListTodoController)

	r.HandleFunc("/team-add", adminteamcontroller.AdminAddTeamController)
	r.HandleFunc("/team-list", adminteamcontroller.AdminListTeamController)

	r.HandleFunc("/team-store", adminteamcontroller.StoreTeamController)

	return r
}
