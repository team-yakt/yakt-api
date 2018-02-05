package main

type Storer interface {
	ListNotes() ([]Note, error)
	ReadNote(id string) (*Note, error)
	WriteNote(n *Note) error
}
