package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewFileStore("markdown")
	data := NewData(store)

	log.Fatal(http.ListenAndServe(":2020", Routes(data)))
}
