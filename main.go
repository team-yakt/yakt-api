package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewStormStore("markdown")
	search := NewBleveSearch("bleve")
	data := NewData(store, search)

	log.Fatal(http.ListenAndServe(":2020", Routes(data)))
}
