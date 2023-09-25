package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Nik      string
	Nama     string
	Password string
}
