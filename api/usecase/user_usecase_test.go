package usecase_test

import (
	"api/domain/model"
	"api/usecase"
	"testing"
)

type mockUserRepository struct {
}

func (t *mockUserRepository) FindByEmail(email string) *model.User {
	return nil
}

func (t *mockUserRepository) SaveUser(user *model.User) (*model.User, error) {
	return &model.User{
		Email: "hoge@gmail.com",
	}, nil
}

func TestSignUpUser(t *testing.T) {
	socialLoginUser := usecase.SocialLoginUser{
		UserID: "hoge",
		Email: "hoge@gmail.com",
		AvatarURL: "http://hoge.example.com",
	}

	mockUserRepository := mockUserRepository{}
	usecase := usecase.NewUserUseCase(&mockUserRepository)

	user, err := usecase.LoginUser(socialLoginUser)
	if err != nil {
		t.Fatalf("login user error")
	}

	if user.Email != "hoge@gmail.com" {
		t.Fatalf("login user error")
	}
}

type mockLoginUserRepository struct {
}

func (t *mockLoginUserRepository) FindByEmail(email string) *model.User {
	return &model.User{
		Email: email,
	}
}

func (t *mockLoginUserRepository) SaveUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func TestLoginUser(t *testing.T) {
	socialLoginUser := usecase.SocialLoginUser{
		UserID: "hoge",
		Email: "hoge@gmail.com",
		AvatarURL: "http://hoge.example.com",
	}

	mockUserRepository := mockUserRepository{}
	usecase := usecase.NewUserUseCase(&mockUserRepository)

	user, err := usecase.LoginUser(socialLoginUser)
	if err != nil {
		t.Fatalf("login user error")
	}

	if user.Email != "hoge@gmail.com" {
		t.Fatalf("login user error")
	}
}


