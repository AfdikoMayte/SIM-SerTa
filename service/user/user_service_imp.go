package service

import (
	"context"
	"database/sql"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/helper"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/domain"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/web"
	repository "github.com/afdikomayte/sim-layanan-sertifikat-tanah/repository/user"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	validator.Validate
}

func (u *UserServiceImpl) Save(ctx context.Context, tx *sql.Tx, request web.UserRequest) web.UserResponse {
	//validation
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, errTx := u.DB.Begin()
	helper.PanicIfError(errTx)
	// handel transaction rollback jika ada err saat insert data
	defer helper.CommitOrRollback(tx)

	//data user yang akan di insert
	userRequest := domain.User{
		Id:       uuid.New(),
		Nik:      request.Nik,
		Nama:     request.Nama,
		Password: request.Password,
	}

	userResult := u.UserRepository.Save(ctx, tx, userRequest)

	return helper.ToUserResponse(userResult)

}

func (u *UserServiceImpl) Update(ctx context.Context, tx *sql.Tx, request web.UserUpdateRequest) web.UserResponse {
	//validation
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, errTx := u.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	//find user to update
	user, err := u.UserRepository.FindById(ctx, tx, request.Id)

	//data user yang di update
	user.Id = request.Id
	user.Nik = request.Nik
	user.Nama = request.Nama
	user.Password = request.Password

	userResult := u.UserRepository.Update(ctx, tx, *user)

	return helper.ToUserResponse(userResult)
}

func (u *UserServiceImpl) Delete(ctx context.Context, tx *sql.Tx, userId uuid.UUID) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//find user to update
	user, err := u.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	u.UserRepository.Delete(ctx, tx, user.Id)

}

func (u *UserServiceImpl) FindById(ctx context.Context, tx *sql.Tx, userId uuid.UUID) web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//find user to update
	user, err := u.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (u *UserServiceImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//find user to update
	users := u.UserRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	//lakukan loop users dari database
	return helper.ToUserResponses(users)
}
