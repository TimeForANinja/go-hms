package sqlUtil

import (
	"database/sql"
	"home_media_server/structs"
)

// GetUsersByCredentials returns a user from database by his pwHash and username
// only returns a single user user is unique
func GetUserByCredentials(db *sql.DB, username string, pwHash string) (*structs.User, error) {
	rows, err := db.Query("SELECT * FROM `users` WHERE pwHash = ? AND name = ?", pwHash, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*structs.User
	for rows.Next() {
		s := structs.User{}
		err := rows.Scan(&s.ID, &s.Name, &s.PwHash, &s.PermissionLevel)
		if err != nil {
			return nil, err
		}
		users = append(users, &s)
	}
	if len(users) != 1 {
		return nil, nil
	}
	return users[0], nil
}

// GetUserByID returns a user from database by his id
func GetUserByID(db *sql.DB, id int) (*structs.User, error) {
	rows, err := db.Query("SELECT * FROM `users` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*structs.User
	for rows.Next() {
		s := structs.User{}
		err := rows.Scan(&s.ID, &s.Name, &s.PwHash, &s.PermissionLevel)
		if err != nil {
			return nil, err
		}
		users = append(users, &s)
	}
	// sth went absolutely wrong when more than one with a UNIQUE id was found
	if len(users) != 1 {
		return nil, nil
	}
	return users[0], nil
}

// AddUser inserts a user into database, user.id will get set after the insert
func AddUser(db *sql.DB, name string, pwHash string, permissionLevel int) (*structs.User, error) {
	s := structs.User{Name: name, PwHash: pwHash, PermissionLevel: permissionLevel}
	return addUser(db, &s)
}

// might export this later on
func addUser(db *sql.DB, user *structs.User) (*structs.User, error) {
	resp, err := db.Exec("INSERT INTO `users` (name, pwHash, permissionLevel) VALUES(?,?,?)", user.Name, user.PwHash, user.PermissionLevel)
	if err != nil {
		return nil, err
	}
	lastInsert, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(lastInsert)
	return user, nil
}
