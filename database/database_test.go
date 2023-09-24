package database_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/database"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/helper"
	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConnetion(t *testing.T) {
	db := database.ConnectDB()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "select * from user")
	helper.PanicIfError(err)

	for rows.Next() {
		var id, nik, nama, password string
		err := rows.Scan(&id, &nik, &nama, &password)
		helper.PanicIfError(err)

		fmt.Println("Id :", id)
		fmt.Println("Nik :", nik)
		fmt.Println("nama :", nama)
		// fmt.Println("Nik :", nik)
	}
	defer rows.Close()

}
