package structs

import "time"

// Download represents a object from downloadQ table WITH foreign items
type Download struct {
	ID     int
	Video  *Video
	Inwork bool
	Failed bool
}

// SimpleDownload represents a object from downloadQ table WITHOUT foreign items
type SimpleDownload struct {
	ID     int
	Video  int
	Inwork bool
	Failed bool
}

// Comment represents a object from comments table WITH foreign items
type Comment struct {
	ID            int
	User          *User
	Video         *Video
	ParentComment *Comment
	Comment       string
	Added         time.Time
}

// SimpleComment represents a object from comments table WITHOUT foreign items
type SimpleComment struct {
	ID            int
	User          int
	Video         int
	ParentComment int
	Comment       string
	Added         time.Time
}

// Rating represents a object from ratings table WITH foreign items
type Rating struct {
	ID     int
	User   *User
	Video  *Video
	Rating int
	Added  time.Time
}

// SimpleRating represents a object from ratings table WITHOUT foreign items
type SimpleRating struct {
	ID     int
	User   int
	Video  int
	Rating int
	Added  time.Time
}
