package web

import "github.com/google/uuid"

type UserResponse struct {
	Id   uuid.UUID
	Nik  string
	Nama string
}
