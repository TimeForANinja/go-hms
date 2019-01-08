package structs

import "time"

type PersonalitySimple struct {
	ID      uint64
	Name    string
	ImdbRef string // nullable
	Created time.Time
	Creator uint64 // nullable => 0
}
type Personality struct {
	ID      uint64
	Name    string
	ImdbRef string // nullable
	Created time.Time
	Creator *User // nullable
}

type LinkPersonalityVideoSimple struct {
	ID          uint64
	Personality uint64
	Video       uint64
}
type LinkPersonalityVideo struct {
	ID          uint64
	Personality *Personality
	Video       *Video
}

type LinkPersonalityMovieSimple struct {
	ID          uint64
	Personality uint64
	Movie       uint64
}
type LinkPersonalityMovie struct {
	ID          uint64
	Personality *Personality
	Movie       *Movie
}
