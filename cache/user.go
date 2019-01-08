package cache

import (
	"errors"

	"github.com/timeforaninja/go-hms/sqlUtil"
	"github.com/timeforaninja/go-hms/structs"
)

var users map[uint64]*structs.User = make(map[uint64]*structs.User)
var usersSimple map[uint64]*structs.UserSimple = make(map[uint64]*structs.UserSimple)

func getUserSimpleByID(id uint64) *structs.UserSimple {
	if _, ok := usersSimple[id]; !ok {
		// key is not present in users
		worked := cacheUserSimpleByID(id)
		if !worked {
			return nil
		}
	}
	return usersSimple[id]
}

func GetUserByID(id uint64) (*structs.User, error) {
	if _, ok := users[id]; !ok {
		err := buildUser(id)
		if err != nil {
			return nil, err
		}
	}
	return users[id], nil
}

func buildUser(id uint64) error {
	userSimple := getUserSimpleByID(id)
	if userSimple == nil {
		return errors.New("no such user")
	}
	user := &structs.User{
		ID:              userSimple.ID,
		Name:            userSimple.Name,
		PwHash:          userSimple.PwHash,
		PermissionLevel: userSimple.PermissionLevel,
		Created:         userSimple.Created,
	}
	users[id] = user
	if userSimple.Creator == 0 {
		return nil
	}
	creator, err := GetUserByID(userSimple.Creator)
	if err != nil {
		uncacheUser(id)
		return err
	}
	user.Creator = creator
	return nil
}

func uncacheUser(id uint64) {
	delete(users, id)
}
func uncacheUserSimple(id uint64) {
	delete(usersSimple, id)
}

func cacheUserSimpleByID(id uint64) bool {
	user, err := sqlUtil.GetUserSimpleByID(db, id)
	if err != nil {
		return false
	}
	usersSimple[user.ID] = user
	return true
}
