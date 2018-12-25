package structs

import "time"

// Download represents a object from downloadQ table
type Download struct {
	ID     int
	Video  *Video
	Inwork bool
	Failed bool
}

// Comment represents a object from comments table
type Comment struct {
	ID            int
	User          *User
	Video         *Video
	ParentComment *Comment
	Comment       string
	Added         time.Time
}

// Rating represents a object from ratings table
type Rating struct {
	ID     int
	User   *User
	Video  *Video
	Rating int
	Added  time.Time
}
