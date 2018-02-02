package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Store interface {
	ListNotes() ([]Note, error)
	ReadNote(id string) (*Note, error)
	WriteNote(n *Note) error
}

type FileStore struct {
	Dir string
}

func NewFileStore(dir string) *FileStore {
	return &FileStore{dir}
}

func (s *FileStore) ListNotes() ([]Note, error) {
	notes := []Note{}
	files, err := ioutil.ReadDir(s.Dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		note, _ := s.ReadNote(f.Name())
		notes = append(notes, *note)
		fmt.Println(filepath.Join(s.Dir, f.Name()))
	}
	return notes, nil
}

func (s *FileStore) ReadNote(filename string) (*Note, error) {
	path := filepath.Join(s.Dir, filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewNoteFromString(string(b))
}

func (s *FileStore) WriteNote(n *Note) error {
	path := filepath.Join(s.Dir, n.Filename())
	return ioutil.WriteFile(path, []byte(n.String()), os.ModePerm)
}
