package main

import (
	"path/filepath"

	"github.com/asdine/storm"
	"github.com/chikamim/pp"
)

type StormStore struct {
	Dir string
}

func NewStormStore(dir string) *StormStore {
	return &StormStore{dir}
}

func (s *StormStore) ListNotes() ([]Note, error) {
	notes := []Note{}
	db, err := s.open()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	err = db.AllByIndex("CreatedAt", &notes, storm.Reverse())
	return notes, err
}

func (s *StormStore) ReadNote(filename string) (*Note, error) {
	db, err := s.open()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	n := &Note{}
	err = db.One("Filename", filename, n)
	pp.Println(n, err)
	return n, err
}

func (s *StormStore) WriteNote(n *Note) error {
	db, err := s.open()
	defer db.Close()
	if err != nil {
		return err
	}
	return db.Save(n)
}

func (s *StormStore) open() (*storm.DB, error) {
	return storm.Open(filepath.Join(s.Dir, "data"))
}
