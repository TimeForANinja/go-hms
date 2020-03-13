package sqlUtil

import (
	"database/sql"
	"fmt"

	"github.com/timeforaninja/go-hms/sqlUtil/sqlQueries"

	// needed for opening sql databases
	_ "github.com/xeodou/go-sqlcipher"
)


// CreateDB creates a new sqlite database file
func CreateDB(file string) *sql.DB {
	fmt.Printf("CreateDB file:\"%s\"\n", file)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(fmt.Errorf("Error \"%s\" while trying to open db", err))
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(fmt.Errorf("Error \"%s\" while enabling foreign keys", err))
	}

	return db
}

// CreateEncryptedDB creates a new encrypted sqlite database file
func CreateEncryptedDB(file string, password string) *sql.DB {
	fmt.Printf("CreateEncryptedDB file:\"%s\" password:\"%s\"\n", file, password)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(fmt.Errorf("Error \"%s\" while trying to open db", err))
	}

	_, err = db.Exec(fmt.Sprintf("PRAGMA key = '%s';", password))
	if err != nil {
		panic(fmt.Errorf("Error \"%s\" while setting password", err))
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(fmt.Errorf("Error \"%s\" while enabling foreign keys", err))
	}

	return db
}
