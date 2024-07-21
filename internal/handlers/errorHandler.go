package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"forum/internal/models"
)

func ErrorHandler(w http.ResponseWriter, code int, st string) {
	tmpl, err := template.ParseFiles("ui/html/pages/error.html")
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(code))
		log.Println("ERROR: " + st + ": " + http.StatusText(code))
		http.Error(w, text, code)
		return
	}
	res := &models.Err{Text_err: http.StatusText(code), Code_err: code}
	err = tmpl.Execute(w, &res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(code))
		log.Println(st + ": " + http.StatusText(code))
		http.Error(w, text, code)
		return
	}
}
