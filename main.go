package main

import (
	"fmt"
	"home_media_server/sqlUtil"
	"home_media_server/structs"
)

func main() {
	db1 := sqlUtil.CreateDB("test2.db")
	defer db1.Close()
	sqlUtil.InitTables(db1)

	db2 := sqlUtil.CreateEncryptedDB("test1.db.enc", "acab")
	defer db2.Close()
	sqlUtil.InitTables(db2)

	// util.StartWeb()

	user := structs.User{Name: "Tobias", PermissionLevel: 9999, PwHash: "1337"}
	sqlUtil.AddUser(db1, &user)
	fmt.Println("----------------")
	sqlUtil.GetUser(db1, "1337")
}
