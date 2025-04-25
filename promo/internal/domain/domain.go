package domain

import (
	"time"
)

type Client struct {
	ID           string    `db:"id"`
	RegisteredAt time.Time `db:"registered_at"`
}

type View struct {
	ID       string    `db:"id"`
	ClientID string    `db:"client_id"`
	EntityID string    `db:"entity_id"`
	ViewedAt time.Time `db:"viewed_at"`
}

type Click struct {
	ID        string    `db:"id"`
	ClientID  string    `db:"client_id"`
	EntityID  string    `db:"entity_id"`
	ClickedAt time.Time `db:"clicked_at"`
}

type Comment struct {
	CommentID   string    `db:"comment_id"`
	ClientID    string    `db:"client_id"`
	EntityID    string    `db:"entity_id"`
	Message     string    `db:"message"`
	CommentedAt time.Time `db:"commented_at"`
}

type Pagination struct {
	Limit  int
	Offset int
}
