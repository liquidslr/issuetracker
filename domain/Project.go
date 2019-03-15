package domain

// Project : The project which are created by the users
type Project struct {
	ID          int64
	Name        string
	OwnerID     int64
	Description string
}

// ProjectService : The  various methods available with a project
type ProjectService interface {
	Project(id int64) (*Project, error)
	Projects() ([]*Project, error)
	Create(p *Project) error
	Delete(id int64) error
}

// ProjectRepository : The  various methods available with an Project
type ProjectRepository interface {
	GetById(id int64) (*Project, error)
	All() ([]*Project, error)
	Create(issue *Project) error
	Delete(id int64) error
}
