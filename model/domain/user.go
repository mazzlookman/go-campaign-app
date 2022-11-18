package domain

import (
	"database/sql"
	"time"
)

type User struct {
	Id             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName sql.NullString
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
