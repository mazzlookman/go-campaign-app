package web

import "database/sql"

type UserFiltered struct {
	Id             int
	Name           string
	Occupation     string
	Email          string
	AvatarFileName sql.NullString
	Role           string
}
