package util

import "time"

// no dependencies
type Categorie struct {
	id              int
	name            string
	parentCategorie int
}
type User struct {
	Id              int
	Name            string
	PwHash          string
	PermissionLevel int
}
type Group struct {
	id              int
	name            string
	permissionLevel int
}

// medium dependencies
type Video struct {
	id              int
	title           string
	downloadRef     string
	equalVideo      *Video
	views           int
	localFile       string
	blockDownload   bool
	blockDownloadBy *User
	added           time.Time
	addedBy         *User
}

// lot of dependencies
type Download struct {
	id     int
	video  *Video
	inwork bool
	failed bool
}
type Comment struct {
	id            int
	user          *User
	video         *Video
	parentComment *Comment
	comment       string
	added         time.Time
}
type Rating struct {
	id     int
	user   *User
	video  *Video
	rating int
	added  time.Time
}
