package database

import (
	"database/sql"
	"time"

	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/helper"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Password#db1@tcp(localhost:3306)/db_sim_layanan_sertifikat_tanah?parseTime=true")
	helper.PanicIfError(err)

	//configurasi pool
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
