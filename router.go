package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(d *Data) chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("list"))
	})
	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("new"))
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	})
	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	return r
}
