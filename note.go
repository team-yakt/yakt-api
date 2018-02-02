package main

import (
	"fmt"
	"strings"
	"time"

	base62 "github.com/mattheath/base62"
	yaml "gopkg.in/yaml.v2"
)

type Note struct {
	ID        string
	Title     string `yaml:"title"`
	Tags      string `yaml:"tags"`
	Body      string
	User      string    `yaml:"user"`
	CreatedAt time.Time `yaml:"created"`
	UpdatedAt time.Time `yaml:"updated"`
}

func NewNote(user, title, tags, body string) *Note {
	n := &Note{}
	n.ID = NewID(time.Now())
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	n.User = user
	n.Title = title
	n.Tags = tags
	n.Body = body

	return n
}

func NewNoteFromString(s string) (*Note, error) {
	l := strings.Split(s, "\n")
	header := strings.Join(l[1:4], "\n")
	body := strings.Join(l[5:], "\n")

	n := &Note{}
	err := yaml.Unmarshal([]byte(header), n)
	if err != nil {
		return nil, err
	}
	n.Body = body
	return n, nil
}

func NewID(t time.Time) string {
	i := t.Local().Sub(t.Truncate(24*time.Hour).Local()) / time.Second
	return base62.EncodeInt64(int64(i))
}

func (n *Note) Filename() string {
	return fmt.Sprintf("%v-%v-%v.md", n.CreatedAt.Local().Format("2006-01-02"), n.User, n.ID)
}

func (n *Note) String() string {
	return fmt.Sprintf(`---
title: %v
tags: %v
user: %v
---
%v
`, n.Title, n.Tags, n.User, n.Body)
}
