package structs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type VideoSimple struct {
	ID              uint64
	Title           string
	DownloadRef     string
	Views           uint64
	LocalFile       string
	blockDownload   bool
	blockDownloadBy uint64
	added           time.Time
	addedBy         uint64
	length          uint64
	resolution      Resolution
	fileHash        string
	fileSize        uint64
}
type Video struct {
	ID              uint64
	Title           string
	DownloadRef     string
	Views           uint64
	LocalFile       string
	blockDownload   bool
	blockDownloadBy *User
	added           time.Time
	addedBy         *User
	length          uint64
	resolution      Resolution
	fileHash        string
	fileSize        uint64
}

type Resolution struct {
	Width, Height uint64
}

func (r Resolution) ToString() string {
	return fmt.Sprintf("%dx%d", r.Width, r.Height)
}
func NewResolution(s string) (Resolution, error) {
	parts := strings.Split(s, "x")
	var r Resolution

	w, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return r, err
	}
	h, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return r, err
	}
	return Resolution{Width: w, Height: h}, nil
}

type LinkIdenticalVideosSimple struct {
	ID     uint64
	VideoA uint64
	VideoB uint64
}
type LinkIdenticalVideos struct {
	ID     uint64
	VideoA *Video
	VideoB *Video
}

type LinkVideoPlaylistSimple struct {
	ID       uint64
	Video    uint64
	Playlist uint64
}
type LinkVideoPlaylist struct {
	ID       uint64
	Video    *Video
	Playlist *Playlist
}
