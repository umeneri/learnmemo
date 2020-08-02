package repository

import "api/domain/model"

type UserRepository interface {
	FindById(id int) model.User
	SetUser(*model.User) error
}

