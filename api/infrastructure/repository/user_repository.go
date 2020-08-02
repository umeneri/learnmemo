package repository

import (
	"api/domain/model"
	"api/domain/repository"

	"github.com/go-xorm/xorm"
)

type userRepository struct {
	dbEngine *xorm.Engine
}

func NewUserRepository(dbName string) repository.UserRepository {
	dbEngine := initDbEngine(dbName)
	return &userRepository{dbEngine}
}

func (t *userRepository) FindByEmail(email string) *model.User {
	return &model.User{}
}

func (t *userRepository) SaveUser(user *model.User) (*model.User, error) {
	_, err := t.dbEngine.Insert(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
