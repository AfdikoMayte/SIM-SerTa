package repository

import (
	"context"
	"database/sql"

	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/domain"
	"github.com/google/uuid"
)

type User interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) *domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) *domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId uuid.UUID)
	FindById(ctx context.Context, tx *sql.Tx, userId uuid.UUID) (*domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []*domain.User
}
