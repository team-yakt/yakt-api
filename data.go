package main

type Data struct {
	Store Store
}

func NewData(s Store) *Data {
	return &Data{s}
}
