package cache

import (
	"database/sql"
	"home_media_server/structs"
)

var db *sql.DB

var groups map[int]*structs.Group = make(map[int]*structs.Group)
var categories map[int]*structs.Categorie = make(map[int]*structs.Categorie)

// DefineDB is used to give the cache access to the db
// it should be the first function being called
func DefineDB(dbRef *sql.DB) {
	db = dbRef
}
