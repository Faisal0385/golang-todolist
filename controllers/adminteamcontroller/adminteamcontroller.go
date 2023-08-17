package adminteamcontroller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
	"todolist/models"

	"golang.org/x/crypto/bcrypt"
)

const uploadDir = "views/static/uploads/teams"

func AdminAddTeamController(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/admin/admin-add-team.html",
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

func AdminListTeamController(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	// Get all matched records
	models.DB.Where("role = ?", "team").Find(&users)

	files := []string{
		"./views/admin/admin-team-list.html",
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
	err = ts.Execute(w, users)

	fmt.Println(users)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func StoreTeamController(w http.ResponseWriter, r *http.Request) {

	full_name := r.FormValue("full_name")
	mobile := r.FormValue("mobile")
	address := r.FormValue("address")
	email := r.FormValue("email")
	password := r.FormValue("password")
	role := "team"
	status := "active"

	// Hash the password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Parse the form data
	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("photo")

	if file != nil {
		if err != nil {
			http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Generate a new unique filename
		newFileName := generateUniqueFileName(handler.Filename)
		newFilePath := filepath.Join(uploadDir, newFileName)
		newFile, err := os.Create(newFilePath)

		if err != nil {
			http.Error(w, "Error creating file", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Error copying file", http.StatusInternalServerError)
			return
		}

		photo := newFilePath

		newUser := models.User{
			FullName: full_name,
			Mobile:   mobile,
			Address:  address,
			Role:     role,
			Email:    email,
			Photo:    photo,
			Password: string(hashedPassword),
			Status:   status,
		}
		models.DB.Create(&newUser)

		// Respond with success message or appropriate status code
		http.Redirect(w, r, "/team-add", http.StatusSeeOther)

	} else {

		newUser := models.User{
			FullName: full_name,
			Mobile:   mobile,
			Address:  address,
			Role:     role,
			Email:    email,
			Password: string(hashedPassword),
			Status:   status,
		}
		models.DB.Create(&newUser)

		// Respond with success message or appropriate status code
		http.Redirect(w, r, "/team-add", http.StatusSeeOther)
	}
}

func generateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	base := filename[:len(filename)-len(ext)]
	return fmt.Sprintf("%s_%d%s", base, uniqueID(), ext)
}

func uniqueID() int64 {
	return time.Now().UnixNano()
}
