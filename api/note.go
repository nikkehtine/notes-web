package api

import (
	"html/template"
	"net/http"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var templs = template.Must(template.ParseFiles("templates/note.html"))

func ReturnNote(w http.ResponseWriter, data *Note) {
	err := templs.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
