package main

import (
	"fmt"
	"home_media_server/sqlUtil"
	"home_media_server/structs"
	"home_media_server/util"
)

func main() {
	db1 := sqlUtil.CreateDB("test2.db")
	defer db1.Close()
	sqlUtil.InitTables(db1)

	db2 := sqlUtil.CreateEncryptedDB("test1.db.enc", "acab")
	defer db2.Close()
	sqlUtil.InitTables(db2)

	go util.StartWeb()

	inputUser := structs.User{Name: "Tobias", PermissionLevel: 9999, PwHash: "1337"}
	sqlUtil.AddUser(db1, &inputUser)
	fmt.Println("----------------")
	sqlUtil.GetUsersByHash(db1, "1337")
	sqlUtil.GetUserByID(db1, 4)
}
