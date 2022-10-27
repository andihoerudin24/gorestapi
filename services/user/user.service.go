package user

import (
	"gorestapi/model/user"
	user2 "gorestapi/repository/user"
)

type ServicesUser interface {
	GetAllUser() *[]user.UserModel
}

type serviceUser struct {
	repository user2.UserRepository
}

func NewServiceUser(repository user2.UserRepository) *serviceUser {
	return &serviceUser{repository: repository}
}

func (r *serviceUser) GetAllUser() *[]user.UserModel {
	datauser := r.repository.GetAllUser()
	if datauser == nil {
		return nil
	}
	return datauser
}
