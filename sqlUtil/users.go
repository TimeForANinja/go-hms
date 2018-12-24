package sqlUtil

import (
	"database/sql"
	"fmt"
	"home_media_server/structs"
)

// GetUser returns a user from database by pwHash
func GetUser(db *sql.DB, pwHash string) *structs.User {
	stmt, err := db.Prepare("SELECT * FROM `users` WHERE pwhash = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pwHash)
	fmt.Println("rows:", rows)
	defer rows.Close()
	var users []*structs.User
	for rows.Next() {
		s := structs.User{}
		rows.Scan(&s.Id, &s.Name, &s.PwHash, &s.PermissionLevel)
		users = append(users, &s)
	}
	for a := 0; a < len(users); a++ {
		fmt.Println(*users[a])
	}
	return users[1]
	//	if size(users) == 0 {
	//		return nil
	//	} else if size(users) > 1 {
	//		return nil
	//	} else {
	//		return users[0]
	//	}
}

// AddUser inserts a user into database
func AddUser(db *sql.DB, user *structs.User) *structs.User {
	stmt, err := db.Prepare("INSERT INTO `users` (name, pwHash, permissionLevel) VALUES(?, ?, ?)")
	fmt.Println(user)
	fmt.Println(*user)
	fmt.Println(stmt)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	resp, err := stmt.Exec(user.Name, user.PwHash, user.PermissionLevel)
	if err != nil {
		panic(err)
	}
	lastInsert, err := resp.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = int(lastInsert)
	return user
}
