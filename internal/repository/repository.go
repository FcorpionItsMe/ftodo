package repository

import "github.com/FcorpionItsMe/ftodo/internal/domain"

type UsersRepository interface {
	SaveUser(inputUser domain.SignUpUserInput) error
	GetUserByLogin(login string) *domain.User
}

type Repository interface {
	UsersRepository
}
