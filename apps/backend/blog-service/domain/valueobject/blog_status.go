package valueobject

// BlogStatus represents the status of a blog post
type BlogStatus string

const (
	// Draft status for unpublished blogs
	Draft BlogStatus = "draft"
	
	// Published status for published blogs
	Published BlogStatus = "published"
	
	// Archived status for archived blogs
	Archived BlogStatus = "archived"
)
