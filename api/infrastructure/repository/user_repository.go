package repository

import (
	"api/domain/model"
	"api/domain/repository"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
)

type userRepository struct {
	dbEngine *xorm.Engine
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{dbEngine}
}

func (t *userRepository) FindByEmail(email string) *model.User {
	user := &model.User{}
	ok, err := t.dbEngine.Table("user").Where("email = ?", email).Get(user)

	if !ok {
		return nil
	}

	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func (t *userRepository) SaveUser(user *model.User) (*model.User, error) {
	_, err := t.dbEngine.Insert(user)
	if err != nil {
		return nil, err
	}
	var newUser model.User
	has, err := t.dbEngine.Table("user").Where("email = ?", user.Email).Get(&newUser)
	if err != nil {
		return nil, err
	}
	if !has {
		err = fmt.Errorf("user cannnot save correctly")
		return nil, err
	}
	return &newUser, err
}

func (t *userRepository) UpdateUser(user *model.User) error {
	_, err := t.dbEngine.Where("email = ?", user.Email).Update(user)
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
