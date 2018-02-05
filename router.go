package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(d *Data) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<a href="/new">new</a>`))
		w.Write([]byte(`<form action="/search/" method="get"><input type="text" name="q"/>
    <input type="submit" value="search" /><pre>`))
		ns, err := d.Store.ListNotes()
		if err != nil {
			RenderErrorHTML(w, r, err)
			return
		}
		RenderNotesHTML(w, r, ns)

		w.Write([]byte(`</pre>`))
	})

	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><body><form action="/note" method="post">
    username:<br/><input type="text" name="user"/><br/>
    title:<br/><input type="text" name="title"/><br/>
		tags:<br/><input type="text" name="tags"/><br/>
		<textarea rows="40" cols="60" name="body"></textarea><br/>
    <input type="submit" value="post">
</form></body></html>`))
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	})

	r.Get("/note/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		n, err := d.Store.ReadNote(id)
		if err != nil {
			RenderErrorHTML(w, r, err)
			return
		}
		RenderNoteHTML(w, r, n)
	})

	r.Post("/note", func(w http.ResponseWriter, r *http.Request) {
		n := &Note{}

		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(n)
			if err != nil {
				w.WriteHeader(500)
				return
			}
		} else {
			r.ParseForm()
			n = NewNote(r.Form["user"][0], r.Form["title"][0], r.Form["tags"][0], r.Form["body"][0])
		}

		err := d.Store.WriteNote(n)
		if err != nil {
			RenderErrorHTML(w, r, err)
			return
		}
		err = d.Search.Index(n)
		if err != nil {
			log.Println("Index failed -", err.Error())
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})

	r.Get("/search/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		q := r.Form["q"][0]
		log.Println("performing query q =", q)
		ids, err := d.Search.Query(q)
		if err != nil {
			RenderErrorHTML(w, r, err)
			return
		}
		for _, id := range ids {
			log.Println("serch res id =", id)
			n, err := d.Store.ReadNote(id)
			if err != nil {
				RenderErrorHTML(w, r, err)
				return
			}
			RenderNoteHTML(w, r, n)
		}
		if len(ids) == 0 {
			w.Write([]byte(`no search result`))
		}
	})

	return r
}
