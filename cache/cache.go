package cache

import (
	"database/sql"
)

var db *sql.DB

// DefineDB is used to give the cache access to the db
// it should be the first function being called
func DefineDB(dbRef *sql.DB) {
	db = dbRef
}

func PopulateCache() error {
	// TODO: write this
	return nil
}
