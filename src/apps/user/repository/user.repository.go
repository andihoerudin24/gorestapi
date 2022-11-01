package repository

import (
	"gorestapi/src/apps/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() *[]model.UserModel
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) *userRepository {
	return &userRepository{connection: connection}
}

func (db *userRepository) GetAllUser() *[]model.UserModel {
	var userModel []model.UserModel
	res := db.connection.Table("users").Where("deleted_at IS NULL").Select("id", "name", "email", "address", "phone").Scan(&userModel)
	if res == nil {
		return nil
	}
	return &userModel
}
