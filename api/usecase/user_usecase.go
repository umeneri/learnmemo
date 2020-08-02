package usecase

import (
	"api/domain/model"
	"api/domain/repository"
)

type SocialLoginUser struct {
	Email     string
	UserID    string
	AvatarURL string
}

type UserUseCase interface {
	LoginUser(user SocialLoginUser) (*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (t *userUseCase) LoginUser(socialLoginUser SocialLoginUser) (*model.User, error) {
	user := t.userRepository.FindByEmail(socialLoginUser.Email)

	if user != nil {
		return user, nil
	}

	user = &model.User{
		ProviderId: socialLoginUser.UserID,
		Email: socialLoginUser.Email,
		AvaterUrl: socialLoginUser.AvatarURL,
	}

	user, err := t.userRepository.SaveUser(user)
	return user, err
}
