package main

import "net/http"

func RenderNoteHTML(w http.ResponseWriter, r *http.Request, n *Note) {
	w.Write([]byte(n.String()))
}

func RenderNotesHTML(w http.ResponseWriter, r *http.Request, ns []Note) {
	for _, n := range ns {
		w.Write([]byte(n.String()))
	}
}

func RenderErrorHTML(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(500)
}

func RenderNoteJSON(w http.ResponseWriter, r *http.Request, n *Note) {
	// TODO:
}

func RenderNotesJSON(w http.ResponseWriter, r *http.Request, ns []Note) {
	// TODO:
}

func RenderErrorJSON(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(500)
}
