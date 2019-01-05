package structs

import "time"

type MovieSimple struct {
	ID              uint64
	Title           string
	ImdbRef         string
	ImdbRating      float64
	LocalFile       string
	BlockDownload   bool
	BlockDownloadBy uint64
	Added           time.Time
	AddedBy         uint64
	Length          uint64
	FileHash        string
	FileSize        uint64
}
type Movie struct {
	ID              uint64
	Title           string
	ImdbRef         string
	ImdbRating      float64
	LocalFile       string
	BlockDownload   bool
	BlockDownloadBy *User
	Added           time.Time
	AddedBy         *User
	Length          uint64
	FileHash        string
	FileSize        uint64
}
