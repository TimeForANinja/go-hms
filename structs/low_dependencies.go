package structs

import "time"

// Categorie represents a object from categories table
type Categorie struct {
	ID              int
	Name            string
	ParentCategorie *Categorie
}

// Video represents a object from videos table
type Video struct {
	ID              int
	Title           string
	DownloadRef     string
	EqualVideo      *Video
	Views           int
	LocalFile       string
	BlockDownload   bool
	BlockDownloadBy *User
	Added           time.Time
	AddedBy         *User
}
