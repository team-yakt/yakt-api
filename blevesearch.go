package main

import (
	"log"
	"os"

	"github.com/blevesearch/bleve"
)

type BleveSearch struct {
	Dir string
}

func NewBleveSearch(dir string) *BleveSearch {
	return &BleveSearch{dir}
}

func (s *BleveSearch) Query(q string) (ids []string, err error) {
	log.Println("bleve query q = ", q)
	index, err := bleve.Open(s.Dir)
	defer index.Close()
	if err != nil {
		return nil, err
	}
	query := bleve.NewQueryStringQuery(q)
	req := bleve.NewSearchRequest(query)
	res, err := index.Search(req)
	if err != nil {
		return nil, err
	}

	for _, hit := range res.Hits {
		ids = append(ids, hit.ID)
	}
	log.Println("bleve query result = ", ids)

	return ids, nil
}

func (s *BleveSearch) Index(n *Note) error {
	if _, err := os.Stat(s.Dir); err == nil {
		indexer, err := bleve.Open(s.Dir)
		defer indexer.Close()
		index(indexer, n)
		return err
	}
	mapping := bleve.NewIndexMapping()
	indexer, err := bleve.New(s.Dir, mapping)
	defer indexer.Close()
	index(indexer, n)
	return err
}

func index(indexer bleve.Index, n *Note) error {
	return indexer.Index(Filename(n), n)
}
