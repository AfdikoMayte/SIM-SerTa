package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/helper"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/domain"
	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) *domain.User {
	//script sql untuk memasukan data ke database
	SQL := "insert into user(id,nik,nama,password) values(?,?,?,?)"

	//prepare Contex
	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)
	//insert with stmt.prepareContex
	_, err = stmt.ExecContext(ctx, user.Id, user.Nik, user.Nama, user.Password)
	helper.PanicIfError(err)
	//close prepare
	defer stmt.Close()

	return &user
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.DB, user domain.User) *domain.User {

	//prepare contex -> execContex
	SQL := "update user set nama = ?, password=? where id = ?"
	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)
	_, err = stmt.ExecContext(ctx, user.Nama, user.Password, user.Id)
	helper.PanicIfError(err)
	defer stmt.Close()

	//return
	return &user
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId uuid.UUID) {
	//prepare content -> execContex
	SQL := "update user set delete_at = now() where id = ?"
	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	_, err = stmt.ExecContext(ctx, userId)
	helper.PanicIfError(err)
	defer stmt.Close()
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId uuid.UUID) (*domain.User, error) {
	//prepare content -> QueryContex
	SQl := "select nik, nama from user where id = ? and deleted_at is null limit 1"
	stmt, err := tx.PrepareContext(ctx, SQl)
	helper.PanicIfError(err)

	rows, err := stmt.QueryContext(ctx, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	//proses data dari query
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Nik, &user.Nama)
		helper.PanicIfError(err)
		return &user, nil
	} else {
		return nil, errors.New("User not found!")
	}
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) *[]domain.User {
	// prepareContex -> queryContex
	SQL := "select nik,nama from user where delete_at is null"
	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)
	rows, err := stmt.QueryContext(ctx)
	helper.PanicIfError(err)

	//users untuk menampung users dari database
	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Nik, &user.Nama)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return &users
}
