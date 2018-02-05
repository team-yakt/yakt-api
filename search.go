package main

type Searcher interface {
	Query(q string) (ids []string, err error)
	Index(n *Note) error
}
