package structs

type DownloadQSimple struct {
	ID     uint64
	Video  uint64
	Inwork bool
	Failed bool
}
type DownloadQ struct {
	ID     uint64
	Video  *Video
	Inwork bool
	Failed bool
}
