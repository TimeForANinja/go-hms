package structs

import "time"

// Categorie represents a object from categories table WITH foreign items
type Categorie struct {
	ID              int
	Name            string
	ParentCategorie *Categorie
}

// SimpleCategorie represents a object from categories table WITHOUT foreign items
type SimpleCategorie struct {
	ID              int
	Name            string
	ParentCategorie int
}

// Video represents a object from videos table WITH foreign items
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

// SimpleVideo represents a object from videos table WITHOUT foreign items
type SimpleVideo struct {
	ID              int
	Title           string
	DownloadRef     string
	EqualVideo      int
	Views           int
	LocalFile       string
	BlockDownload   bool
	BlockDownloadBy int
	Added           time.Time
	AddedBy         int
}
