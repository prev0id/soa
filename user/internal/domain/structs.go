package domain

import (
	"time"
)

type User struct {
	ID        int64
	Login     string
	Email     string
	FirstName string
	LastName  string
	BirthDate time.Time
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
