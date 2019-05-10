package main

import (
	"github.com/timeforaninja/go-hms/sqlUtil"
)

func main() {
	db := sqlUtil.CreateDB("test.db")
	defer db.Close()
	sqlUtil.InitTables(db)
}
