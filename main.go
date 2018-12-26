package main

import (
	"fmt"
	"home_media_server/cache"
	"home_media_server/sqlUtil"
	"home_media_server/util"
)

func main() {
	db1 := sqlUtil.CreateDB("test2.db")
	defer db1.Close()
	sqlUtil.InitTables(db1)

	db2 := sqlUtil.CreateEncryptedDB("test1.db.enc", "acab")
	defer db2.Close()
	sqlUtil.InitTables(db2)

	cache.DefineDB(db1)

	go util.StartWeb()

	// Testing:
	sqlUtil.AddUser(db1, "Tobias", "1337", 9999)
	fmt.Println("----------------")
	sqlUtil.GetUserByCredentials(db1, "Tobias", "1337")
	sqlUtil.GetUserByID(db1, 4)
}
