package sqlUtil

import (
	"database/sql"
	"errors"
	"time"

	"github.com/timeforaninja/go-hms/structs"
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
		// TODO: missing some variables
		// TODO: created is not in format time
		// TODO: check what happens when Creator = null
		err := rows.Scan(&s.ID, &s.Name, &s.PwHash, &s.PermissionLevel)
		if err != nil {
			return nil, err
		}
		users = append(users, &s)
	}
	if len(users) != 1 {
		return nil, errors.New("not exactly one user found")
	}
	return users[0], nil
}

// GetUserByID returns a user from database by his id
func GetUserSimpleByID(db *sql.DB, id uint64) (*structs.UserSimple, error) {
	rows, err := db.Query("SELECT * FROM `users` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*structs.UserSimple
	for rows.Next() {
		s := structs.UserSimple{}
		// TODO: created is not in format time
		// TODO: check what happens when Creator = null
		err := rows.Scan(&s.ID, &s.Name, &s.PwHash, &s.PermissionLevel, &s.Created, &s.Creator)
		if err != nil {
			return nil, err
		}
		users = append(users, &s)
	}
	// sth went absolutely wrong when more than one with a UNIQUE id was found
	if len(users) != 1 {
		return nil, errors.New("not exactly one user found")
	}
	return users[0], nil
}

// AddUser inserts a user into database, user.id will get set after the insert
func AddUser(db *sql.DB, name string, pwHash string, permissionLevel uint64, creator *structs.User) (*structs.User, error) {
	s := structs.User{
		Name:            name,
		PwHash:          pwHash,
		PermissionLevel: permissionLevel,
		Created:         time.Now(),
		Creator:         creator,
	}
	return addUser(db, &s)
}

// might export this later on
func addUser(db *sql.DB, user *structs.User) (*structs.User, error) {
	resp, err := db.Exec("INSERT INTO `users` (name, pwHash, permissionLevel, created, creator) VALUES(?,?,?,?,?)", user.Name, user.PwHash, user.PermissionLevel, user.Created.Unix(), user.Creator.ID)
	if err != nil {
		return nil, err
	}
	lastInsert, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = uint64(lastInsert)
	return user, nil
}
