package sqlUtil

import (
	"database/sql"
	"fmt"

	"github.com/timeforaninja/go-hms/sqlUtil/sqlQueries"

	// needed for opening sql databases
	_ "github.com/xeodou/go-sqlcipher"
)

// InitTables creates default tables in a database
func InitTables(db *sql.DB) {
	querys := []string{
		sqlQueries.QueryCreateConfig, // config.go

		sqlQueries.QueryCreateCategories, // categories.go

		sqlQueries.QueryCreateLinkCategorieUserPermission,  // categories.go
		sqlQueries.QueryCreateLinkCategoriePersonality,     // categories.go
		sqlQueries.QueryCreateLinkCategorieVideo,           // categories.go
		sqlQueries.QueryCreateLinkCategorieMovie,           // categories.go
		sqlQueries.QueryCreateLinkCategorieGroupPermission, // categories.go
		sqlQueries.QueryCreateLinkCategoriePlaylist,        // categories.go

		sqlQueries.QueryCreateUsers,                // users.go
		sqlQueries.QueryCreateGroups,               // groups.go
		sqlQueries.QueryCreateMovies,               // movies.go
		sqlQueries.QueryCreateVideos,               // videos.go
		sqlQueries.QueryCreateDownloadQ,            // downloadQ.go
		sqlQueries.QueryCreatePersonalities,        // personalities.go
		sqlQueries.QueryCreateLinkPersonalityVideo, // personalities.go
		sqlQueries.QueryCreateLinkPersonalityMovie, // personalities.go
		sqlQueries.QueryCreatePlaylists,            // playlists.go
		sqlQueries.QueryCreateLinkUserGroup,        // users.go
		sqlQueries.QueryCreateComments,             // users.go
		sqlQueries.QueryCreateRatings,              // users.go
		sqlQueries.QueryCreateLinkIdenticalVideos,  // videos.go
		sqlQueries.QueryCreateLinkVideoPlaylist,    // videos.go

		sqlQueries.QueryDefaultAdminUser,       // defaults.go
		sqlQueries.QueryDefaultEveryoneGroup,   // defaults.go
		sqlQueries.QueryDefaultGroup,           // defaults.go
		sqlQueries.QueryDefaultAdminUserGroups, // defaults.go
	}
	for i := 0; i < len(querys); i++ {
		query := querys[i]
		_, err := db.Exec(query)
		if err != nil {
			panic(fmt.Errorf("Error \"%s\" running query:\"%s\"", err, query))
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
