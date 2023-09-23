package api

import (
	"bytes"
	"log"
	"net/http"
)

var buf bytes.Buffer

var db = []Note{
	{
		ID:      1,
		Title:   "Hello World",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		ID:      2,
		Title:   "Do some magic",
		Content: "Edit this note or delete the whole thing!",
	},
}

// ReturnAllNotes is a function that writes all the notes to the given http.ResponseWriter.
func ReturnAllNotes(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(db); i++ {
		err := templs.Execute(&buf, db[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}
