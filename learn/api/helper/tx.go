package helper

import (
	"database/sql"
	"learn/api/helper"
)

func CommitOrRollback(tx *sql.Tx) {
	err2 := recover()
	if err2 != nil {
		errRollback := tx.Rollback()
		helper.PanicIfError(errRollback)
		panic(err2)
	} else {
		errCommit := tx.Commit()
		helper.PanicIfError(errCommit)
	}
}
