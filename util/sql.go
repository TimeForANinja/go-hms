package util

import (
	"database/sql"
	"fmt"

	// needed for opening sql databases
	_ "github.com/xeodou/go-sqlcipher"
)

// InitTables creates default tables in a database
func InitTables(db *sql.DB) {

	// static creation querys
	querys := [7]string{
		"CREATE TABLE IF NOT EXISTS `downloadQ` (`id` INTEGER PRIMARY KEY)",
		"CREATE TABLE IF NOT EXISTS `videos` (`id` INTEGER PRIMARY KEY)",
		"CREATE TABLE IF NOT EXISTS `categories` (`id` INTEGER PRIMARY KEY)",
		"CREATE TABLE IF NOT EXISTS `comments` (`id` INTEGER PRIMARY KEY)",
		"CREATE TABLE IF NOT EXISTS `ratings` (`id` INTEGER PRIMARY KEY)",
		"CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY)",
		"CREATE TABLE IF NOT EXISTS `groups` (`id` INTEGER PRIMARY KEY)",
	}
	for i := 0; i < len(querys); i++ {
		query := querys[i]
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Error \"%s\" executing query:\"%s\"", err, query)
			return
		}
	}
}

// CreateDB creates a new sqlite database file
func CreateDB(file string) *sql.DB {
	fmt.Printf("CreateDB file:\"%s\"\n", file)

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println("err1", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println("err2", err)
	}

	return db
}

// CreateEncryptedDB creates a new encrypted sqlite database file
func CreateEncryptedDB(file string, password string) *sql.DB {
	fmt.Printf("CreateEncryptedDB file:\"%s\" password:\"%s\"\n", file, password)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println("err1", err)
	}

	p := fmt.Sprintf("PRAGMA key = '%s';", password)
	_, err = db.Exec(p)
	if err != nil {
		fmt.Println("err2", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println("err3", err)
	}

	return db
}
