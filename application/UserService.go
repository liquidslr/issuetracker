package application

import "github.com/liquidslr/issuetracker/domain"

type UserService struct {
	UsersRepository domain.UserRepository
}

func (s UserService) Users() ([]*domain.User, error) {
	return s.UsersRepository.All()
}

func (s UserService) Create(u *domain.User) error {
	return s.UsersRepository.Create(u)
}

func (s UserService) Delete(id int64) error {
	return s.UsersRepository.Delete(id)
}

func (s UserService) User(id int64) (*domain.User, error) {
	return s.UsersRepository.GetById(id)
}
