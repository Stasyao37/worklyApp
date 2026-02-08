package models

import (
	"time"
)

type NotesModel struct {
	notes   []Note
	cursor  int
	filter  string
	editing bool
}

type Note struct {
	ID        int
	Title     string
	Content   string
	Tags      bool
	CreatedAt time.Time
	Pinned    bool
}
