package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Nik      int
	Nama     string
	Password string
}
