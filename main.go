package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var users = make(map[string]*User)
var tmpl = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/login.html",
	"templates/register.html",
))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, ok := users[email]
	if !ok {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !checkPassword(password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", email)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "register.html", nil)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if _, exists := users[email]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := &User{Email: email, Password: hashedPassword}
	users[email] = user

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

