package web

import "github.com/google/uuid"

type UserRequest struct {
	Id       uuid.UUID
	Nik      string
	Nama     string
	Password string
}
