package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if recover() != nil {
		err := tx.Rollback()
		PanicIfError(err)
	} else {
		err := tx.Commit()
		PanicIfError(err)
	}
}
