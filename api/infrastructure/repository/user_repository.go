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
	user := &model.User{}
	ok, err := t.dbEngine.Table("user").Where("email = ?", email).Get(user)

	if !ok {
		return nil
	}

	if err != nil {
		panic(err)
	}
	return user
}

func (t *userRepository) SaveUser(user *model.User) (*model.User, error) {
	_, err := t.dbEngine.Insert(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *userRepository) UpdateUser(user *model.User) error {
	_, err := t.dbEngine.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (t *userRepository) DeleteUser(user *model.User) error {
	_, err := t.dbEngine.Id(user.Id).Delete(model.User{})
	if err != nil {
		return err
	}
	return nil
}
