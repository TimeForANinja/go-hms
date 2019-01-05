package structs

import "time"

type UserSimple struct {
	ID              uint64
	Name            string
	PwHash          string
	PermissionLevel uint64
	Created         time.Time
	Creator         uint64
}
type User struct {
	ID              uint64
	Name            string
	PwHash          string
	PermissionLevel uint64
	Created         time.Time
	Creator         *User
}

type LinkUserGroupSimple struct {
	ID    uint64
	User  uint64
	Group uint64
}
type LinkUserGroup struct {
	ID    uint64
	User  *User
	Group *Group
}

type CommentSimple struct {
	ID            uint64
	User          uint64
	Video         uint64
	ParentComment uint64
	Comment       string
	added         time.Time
}
type Comment struct {
	ID            uint64
	User          *User
	Video         *Video
	ParentComment *Comment
	Comment       string
	added         time.Time
}

type RatingSimple struct {
	ID     uint64
	User   uint64
	Video  uint64
	Rating uint8
	Added  time.Time
}
type Rating struct {
	ID     uint64
	User   *User
	Video  *Video
	Rating uint8
	Added  time.Time
}
