package main

import (
	"fmt"
	"time"
)

func ExampleNoteString() {
	note := NewNote("user", "title", "tags", "body")
	fmt.Println(note.String())
	// Output:
	// ---
	// title: title
	// tags: tags
	// user: user
	// ---
	// body
}

func ExampleNewNoteString() {
	s := `---
title: title
tags: tags
user: user
---
body`
	note, _ := NewNoteFromString(s)
	fmt.Println(note.String())
	// Output:
	// ---
	// title: title
	// tags: tags
	// user: user
	// ---
	// body
}

func ExampleNoteFilename() {
	note := NewNote("user", "title", "tags", "body")
	note.ID = "123456"
	note.CreatedAt = time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
	fmt.Println(note.Filename())
	// Output: 2014-12-31-user-123456.md
}

func ExampleNewID() {
	id := NewID(time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC))
	fmt.Println(id)
	// Output: BRk
}
