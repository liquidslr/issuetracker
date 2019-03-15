package domain

// User : The User which are created by the users
type User struct {
	ID      int64
	Name    string
	SurName string
	Email   string
}

// UserService : The  various methods available with an User
type UserService interface {
	User(id int64) (*User, error)
	Users() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}

// UserRepository : The  various methods available with an User
type UserRepository interface {
	GetById(id int64) (*User, error)
	All() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}
