package domain

// Issue : The issue which are created by the users
type Issue struct {
	ID          int64
	Title       string
	Description string
	ProjectID   int64
	OwnerID     int64
	Status      string
}
