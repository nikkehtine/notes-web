package api

import (
	"html/template"
	"net/http"
	"strconv"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var templs = template.Must(template.ParseFiles("templates/note.html"))

func ReturnNote(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Note ID is missing", http.StatusBadRequest)
		return
	}

	noteID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note := db[noteID]

	err = templs.Execute(w, note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
