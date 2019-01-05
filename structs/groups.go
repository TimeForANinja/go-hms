package structs

import "time"

type GroupSimple struct {
	ID              uint64
	Name            string
	PermissionLevel uint64
	Created         time.Time
	Creator         int
}
type Group struct {
	ID              uint64
	Name            string
	PermissionLevel uint64
	Created         time.Time
	Creator         *User
}
