package structs

import "time"

type MovieSimple struct {
	ID              uint64
	Title           string // TODO: format
	AltTitles       string
	ImdbRef         string
	ImdbRating      float64
	Trailer         uint64 // nullable => 0
	CoverFile       string // nullable
	LocalFile       string // nullable
	BlockDownload   bool
	BlockDownloadBy uint64 // nullable => 0
	Added           time.Time
	AddedBy         uint64 // nullable => 0
	Length          uint64
	FileHash        string // nullable
	FileSize        uint64 // nullable => 0
}
type Movie struct {
	ID              uint64
	Title           string // TODO: format
	AltTitles       string
	ImdbRef         string
	ImdbRating      float64
	Trailer         *Video // nullable
	CoverFile       string // nullable
	LocalFile       string // nullable
	BlockDownload   bool
	BlockDownloadBy *User // nullable
	Added           time.Time
	AddedBy         *User // nullable
	Length          uint64
	FileHash        string // nullable
	FileSize        uint64 // nullable => 0
}
