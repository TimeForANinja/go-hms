package structs

import "time"

type PlaylistSimple struct {
	ID      uint64
	Name    string
	Public  bool
	Created time.Time
	Creator uint64 // nullable => 0
}
type Playlist struct {
	ID      uint64
	Name    string
	Public  bool
	Created time.Time
	Creator *User // nullable
}
