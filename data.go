package main

type Data struct {
	Store  Storer
	Search Searcher
}

func NewData(store Storer, search Searcher) *Data {
	return &Data{store, search}
}
