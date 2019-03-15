package domain

// Issue : The issue which are created by the users
type Issue struct {
	ID          int64
	Title       string
	Description string
	ProjectID   int64
	OwnerID     int64
	Status      string
	Priority    Priority
}

// IssueService : The  various methods available with an issue
type IssueService interface {
	Issue(id int64) (*Issue, error)
	Issues() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}

// IssueRepository : The  various methods available with an issue
type IssueRepository interface {
	GetById(id int64) (*Issue, error)
	All() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}
