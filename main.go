package main

import (
	Util "home_media_server/util"
)

func main() {
	db1 := Util.CreateDB("test2.db")
	defer db1.Close()
	Util.InitTables(db1)

	db2 := Util.CreateEncryptedDB("test1.db.enc", "acab")
	defer db2.Close()
	Util.InitTables(db2)

	Util.StartWeb()
}
