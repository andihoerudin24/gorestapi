package service

import (
	"gorestapi/src/apps/user/model"
	"gorestapi/src/apps/user/repository"
)

type UserService interface {
	GetAllUser(perPage int64, offsets int64) (*[]model.UserModel, int64)
	CreateUser(userModel model.UserModel) (*model.UserModel, error)
	FindById(id int64) (*model.UserModel, error)
	Update(id int64, userModel model.UserModel) int64
	Delete(id int64) error
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

func (u *userService) FindById(id int64) (*model.UserModel, error) {
	FindUser, err := u.userRepository.FindById(id)
	return FindUser, err
}

func (u *userService) Update(id int64, userModel model.UserModel) int64 {
	RowsAffected := u.userRepository.Update(id, userModel)
	return RowsAffected
}

func (u *userService) Delete(id int64) error {
	Delete := u.userRepository.Delete(id)
	return Delete
}
