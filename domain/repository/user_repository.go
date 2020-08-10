package repository

import "api/domain/model"

type UserRepository interface {
	FindByEmail(email string) *model.User
	SaveUser(*model.User) (*model.User, error)
	UpdateUser(*model.User) error
	DeleteUser(*model.User) error
}

