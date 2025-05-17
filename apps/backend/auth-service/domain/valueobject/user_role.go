package valueobject

// UserRole represents the role of a user
type UserRole string

const (
	// RoleUser is the standard user role
	RoleUser UserRole = "user"
	
	// RoleAdmin is the administrator role
	RoleAdmin UserRole = "admin"
	
	// RoleAuthor is the author role
	RoleAuthor UserRole = "author"
)
