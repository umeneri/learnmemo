package usecase

import (
	"api/domain/model"
	"api/domain/repository"
	"time"
)

type SocialLoginUser struct {
	Email     string
	UserID    string
	AvatarURL string
}

type UserUseCase interface {
	SaveUser(user SocialLoginUser) (*model.User, error)
	FindByEmail(email string) *model.User
	UpdateUser(user *model.User) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (t *userUseCase) SaveUser(socialLoginUser SocialLoginUser) (*model.User, error) {
	user := t.userRepository.FindByEmail(socialLoginUser.Email)

	if user != nil {
		return user, nil
	}

	user = &model.User{
		Email:      socialLoginUser.Email,
		Name:       socialLoginUser.UserID,
		ProviderId: socialLoginUser.UserID,
		AvatarUrl:  socialLoginUser.AvatarURL,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	user, err := t.userRepository.SaveUser(user)
	return user, err
}

func (t *userUseCase) FindByEmail(email string) *model.User {
	 return t.userRepository.FindByEmail(email)
}

func (t *userUseCase) UpdateUser(user *model.User) error {
		err := t.userRepository.UpdateUser(user)
		return err
}