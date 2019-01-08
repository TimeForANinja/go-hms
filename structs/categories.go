package structs

type CategorieSimple struct {
	ID              uint64
	Name            string
	ParentCategorie uint64 // nullable => 0
}
type Categorie struct {
	ID              uint64
	Name            string
	ParentCategorie *Categorie // nullable
}

type LinkCategorieUserPermissionSimple struct {
	ID        uint64
	User      uint64
	Categorie uint64
	CanRead   bool
	CanWrite  bool
}
type LinkCategorieUserPermission struct {
	ID        uint64
	User      *User
	Categorie *Categorie
	CanRead   bool
	CanWrite  bool
}

type LinkCategoriePersonalitySimple struct {
	ID          uint64
	Personality uint64
	Categorie   uint64
}
type LinkCategoriePersonality struct {
	ID          uint64
	Personality *Personality
	Categorie   *Categorie
}

type LinkCategorieVideoSimple struct {
	ID        uint64
	Video     uint64
	Categorie uint64
}
type LinkCategorieVideo struct {
	ID        uint64
	Video     *Video
	Categorie *Categorie
}

type LinkCategorieMovieSimple struct {
	ID        uint64
	Movie     uint64
	Categorie uint64
}
type LinkCategorieMovie struct {
	ID        uint64
	Movie     *Movie
	Categorie *Categorie
}

type LinkCategorieGroupPermissionSimple struct {
	ID        uint64
	Group     uint64
	Categorie uint64
	CanRead   bool
	CanWrite  bool
}
type LinkCategorieGroupPermission struct {
	ID        uint64
	Group     *Group
	Categorie *Categorie
	CanRead   bool
	CanWrite  bool
}

type LinkCategoriePlaylistSimple struct {
	ID        uint64
	Playlist  uint64
	Categorie uint64
}
type LinkCategoriePlaylist struct {
	ID        uint64
	Playlist  *Playlist
	Categorie *Categorie
}
