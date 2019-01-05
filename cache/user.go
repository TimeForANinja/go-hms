package cache

import (
	"home_media_server/sqlUtil"
	"home_media_server/structs"
)

var users map[int]*structs.User = make(map[int]*structs.User)

func getUserByID(id int) *structs.User {
	if _, ok := users[id]; !ok {
		// key is not present in users
		worked := cacheUserByID(id)
		if !worked {
			return nil
		}
	}
	return users[id]
}

func uncacheUser(id int) {
	delete(users, id)
}

func cacheUserByID(id int) bool {
	user, err := sqlUtil.GetUserByID(db, id)
	if err != nil || user == nil {
		return false
	}
	users[user.ID] = user
	return true
}