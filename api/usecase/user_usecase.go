package usecase

import (
	"api/domain/repository"
)

type SocialLoginUser struct {
	Email     string
	UserID    string
	AvatarURL string
}

type UserUseCase interface {
	LoginUser(user SocialLoginUser) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (t *userUseCase) LoginUser(user SocialLoginUser) error {
	// return t.userRepository.(User)
	return nil
}
