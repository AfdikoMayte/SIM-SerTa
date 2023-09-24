package helper

import "database/sql"

func commitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		tx.Rollback()
	}
}
