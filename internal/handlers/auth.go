package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"registration-app/internal/db"

	"golang.org/x/crypto/bcrypt"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		var exists bool
		error := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
		if error != nil || exists {
			templates.ExecuteTemplate(w, "registration.html", map[string]string{
				"Error": "User already registered with this email.",
			})
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		_, err = db.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
			username, email, string(hashed))
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	templates.ExecuteTemplate(w, "registration.html", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// var username, hashed string
		// err := db.DB.QueryRow("SELECT username, password FROM users WHERE email=$1", email).Scan(&username, &hashed)
		// if err != nil {
		// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		// 	return
		// }
		var userID int
		var hashedPassword string

		// Query the database for user details
		err := db.DB.QueryRow("SELECT id, password FROM users WHERE email=$1", email).Scan(&userID, &hashedPassword)
		if err != nil {
			// Invalid credentials
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "session_user",
			Value: strconv.Itoa(userID),
			Path:  "/",
		})
		http.Redirect(w, r, "/shop", http.StatusSeeOther)
		return
	}

	templates.ExecuteTemplate(w, "login.html", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session_user",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
