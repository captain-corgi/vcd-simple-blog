package valueobject

// ProfileStatus represents the status of a user profile
type ProfileStatus string

const (
	// ProfileStatusPublic indicates a public profile
	ProfileStatusPublic ProfileStatus = "public"

	// ProfileStatusPrivate indicates a private profile
	ProfileStatusPrivate ProfileStatus = "private"

	// ProfileStatusLimited indicates a limited visibility profile
	ProfileStatusLimited ProfileStatus = "limited"
)
