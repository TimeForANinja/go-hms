package structs

// User represents a object from users table
type User struct {
	ID              int
	Name            string
	PwHash          string
	PermissionLevel int
}

// Group represents a object from groups table
type Group struct {
	ID              int
	Name            string
	PermissionLevel int
}
