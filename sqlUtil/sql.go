package sqlUtil

import (
	"database/sql"
	"fmt"
	"home_media_server/sqlUtil/sqlQueries"
	// needed for opening sql databases
	_ "github.com/xeodou/go-sqlcipher"
)

// InitTables creates default tables in a database
func InitTables(db *sql.DB) {
	querys := []string{
		// no dependencies
		sqlQueries.QueryCreateCategories,
		sqlQueries.QueryCreateUsers,
		sqlQueries.QueryCreateGroups,
		// low dependencies
		sqlQueries.QueryCreateVideos,
		// medium dependencies
		sqlQueries.QueryCreateDownloadQ,
		sqlQueries.QueryCreateComments,
		sqlQueries.QueryCreateRatings,
		// high dependencies
		sqlQueries.QueryCreateLinksUserGroup,
		sqlQueries.QueryCreateLinksVideoCategorie,
		sqlQueries.QueryCreateLinksCategorieGroupPermission,
		sqlQueries.QueryCreateLinksCategorieUserPermission,
	}
	for i := 0; i < len(querys); i++ {
		query := querys[i]
		_, err := db.Exec(query)
		if err != nil {
			panic(fmt.Errorf("Error \"%s\" creating a table with query:\"%s\"", err, query))
		}
	}
}

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
