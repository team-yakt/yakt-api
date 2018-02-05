package main

type Store interface {
	ListNotes() ([]Note, error)
	ReadNote(id string) (*Note, error)
	WriteNote(n *Note) error
}
