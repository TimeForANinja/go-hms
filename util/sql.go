package util

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	// needed for opening sql databases
	_ "github.com/xeodou/go-sqlcipher"
)

// InitTables creates default tables in a database
func InitTables(db *sql.DB) {
	absPath, err := filepath.Abs("./util/querys.sql")
	if err != nil {
		panic(fmt.Errorf("Error getting filepath: \"%s\"", err))
	}
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(fmt.Errorf("Error loading querys: \"%s\"", err))
	}
	querys := strings.Split(string(data), ");\n")
	for i := 0; i < len(querys)-1; i++ {
		query := querys[i] + ");"
		_, err := db.Exec(query)
		if err != nil {
			panic(fmt.Errorf("Error \"%s\" executing query:\"%s\"", err, query))
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
