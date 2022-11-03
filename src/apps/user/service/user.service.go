package service

import (
	"gorestapi/src/apps/user/model"
	"gorestapi/src/apps/user/repository"
)

type UserService interface {
	GetAllUser(perPage int64, offsets int64) (*[]model.UserModel, int64)
	CreateUser(userModel model.UserModel) (*model.UserModel, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (u *userService) GetAllUser(perPage int64, offsets int64) (*[]model.UserModel, int64) {
	dataUser, totalrows := u.userRepository.GetAllUser(perPage, offsets)
	if dataUser == nil {
		return nil, totalrows
	}
	return dataUser, totalrows
}

func (u *userService) CreateUser(userModel model.UserModel) (*model.UserModel, error) {
	createUser, err := u.userRepository.CreateUser(userModel)
	return createUser, err
}
