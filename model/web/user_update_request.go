package web

import "github.com/google/uuid"

type UserUpdateRequest struct {
	Id       uuid.UUID
	Nik      string `validate:"required, len=16"`
	Nama     string `validate:"required, min=3, max=100"`
	Password string
}
