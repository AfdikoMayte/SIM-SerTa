package service

import (
	"context"
	"database/sql"

	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/helper"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/web"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/repository"
	"github.com/google/uuid"
)

type UserImpl struct {
	UserRepository repository.User
	DB             *sql.DB
}

func (service *UserImpl) Create(ctx context.Context, request web.UserRequest) *web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer 
}

func (service *UserImpl) Update(ctx context.Context, request web.UserRequest) *web.UserResponse {

}

func (service *UserImpl) Destroy(ctx context.Context, userId uuid.UUID) {

}

func (service *UserImpl) FindById(ctx context.Context, userId uuid.UUID) *web.UserResponse {

}

func (service *UserImpl) FindAll(ctx context.Context) []*web.UserResponse {

}
