package usecase

import (
	"api/domain/repository"
)

type UserUseCase interface {
	LoginUser() error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
  return &userUseCase{
		userRepository: userRepository,
	}
}

func (t *userUseCase) LoginUser() error {
	// return t.userRepository.(User)
	return nil
}