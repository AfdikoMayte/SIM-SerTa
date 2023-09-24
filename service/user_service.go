package service

import (
	"context"

	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/web"
	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, request web.UserRequest) *web.UserResponse
	Update(ctx context.Context, request web.UserRequest) *web.UserResponse
	Destroy(ctx context.Context, userId uuid.UUID)
	FindById(ctx context.Context)
	FindAll(ctx context.Context)
}
