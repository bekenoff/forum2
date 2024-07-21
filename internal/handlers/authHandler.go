package handlers

import (
	"fmt"
	"forum/internal/models"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) CreateClient(w http.ResponseWriter, r *http.Request) {

	nameFunction := "Register"
	if r.URL.Path != "/register" {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}
	tmpl, err := template.ParseFiles("/Users/bekenov/Desktop/forum-main/web/html/pages/signup.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, nameFunction)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, nil); err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}
	case "POST":
		user := &models.UserRegistration{
			Username: r.FormValue("name"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		fmt.Print(user)

		if err := h.service.CreateClient(*user); err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			log.Fatal(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, nameFunction)
	}
}
