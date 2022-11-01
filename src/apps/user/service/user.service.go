package service

import (
	"gorestapi/src/apps/user/model"
	"gorestapi/src/apps/user/repository"
)

type UserService interface {
	GetAllUser() *[]model.UserModel
	CreateUser(userModel model.UserModel) (model.UserModel, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (u *userService) GetAllUser() *[]model.UserModel {
	dataUser := u.userRepository.GetAllUser()
	if dataUser == nil {
		return nil
	}
	return dataUser
}

func (u *userService) CreateUser(userModel model.UserModel) (model.UserModel, error) {
	createUser, err := u.userRepository.CreateUser(userModel)
	return createUser, err
}
