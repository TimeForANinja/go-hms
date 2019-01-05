package structs

import "time"

type PlaylistSimple struct {
	ID        uint64
	Name      string
	AltTitles string
	Public    bool
	Created   time.Time
	Creator   uint64
}
type Playlist struct {
	ID        uint64
	Name      string
	AltTitles string
	Public    bool
	Created   time.Time
	Creator   *User
}
