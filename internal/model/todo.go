package model

import (
	"database/sql"
	"time"
)

type TodoNote struct {
	ID        int64
	Info      TodoNoteInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type TodoNoteInfo struct {
	Title   string
	Content string
}
