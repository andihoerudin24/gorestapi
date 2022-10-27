package user

import (
	"gorestapi/model/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() *[]user.UserModel
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) *userRepository {
	return &userRepository{connection: connection}
}

func (db *userRepository) GetAllUser() *[]user.UserModel {
	var userModel []user.UserModel
	res := db.connection.Table("users").Select("id", "name", "email", "address", "phone").Scan(&userModel)
	if res == nil {
		return nil
	}
	return &userModel
}
