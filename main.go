package main

import (
	"github.com/timeforaninja/go-hms/cache"
	"github.com/timeforaninja/go-hms/sqlUtil"
)

func main() {
	db1 := sqlUtil.CreateDB("test.db")
	defer db1.Close()
	sqlUtil.InitTables(db1)

	db2 := sqlUtil.CreateEncryptedDB("test.db.enc", "1337")
	defer db2.Close()
	sqlUtil.InitTables(db2)

	cache.DefineDB(db1)
	cache.PopulateCache()

	// Testing:
	//	fmt.Println(sqlUtil.AddUser(db1, "Tobias", "1337", 9999))
	//	fmt.Println("----------------")
	//	fmt.Println(sqlUtil.GetUserByCredentials(db1, "Tobias", "1337"))
	//	fmt.Println(sqlUtil.GetUserByID(db1, 4))

	//	util.StartWeb()
}
