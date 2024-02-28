// main.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "projetoGo"
)

type User struct {
	Name  string
	Email string
	Phone string
}

type PageVariables struct {
	Title string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexão com o PostgreSQL estabelecida com sucesso")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		newUser := User{
			Name:  name,
			Email: email,
			Phone: phone,
		}

		err := insertUser(newUser)
		if err != nil {
			http.Error(w, "Erro ao inserir usuário no banco de dados", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/success", http.StatusSeeOther)
		return
	}

	data := PageVariables{
		Title: "Cadastro de Usuários:",
	}

	renderTemplate(w, "form.html", data)
}

func insertUser(user User) error {
	_, err := db.Exec("INSERT INTO users (name, email, phone) VALUES ($1, $2, $3)", user.Name, user.Email, user.Phone)
	return err
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cadastro realizado com sucesso!")
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageVariables) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/success", successHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
