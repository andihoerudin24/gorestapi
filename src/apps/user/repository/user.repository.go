package repository

import (
	"gorestapi/src/apps/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(perPage int64, offset int64) (*[]model.UserModel, int64)
	CreateUser(userModel model.UserModel) (*model.UserModel, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) *userRepository {
	return &userRepository{connection: connection}
}

func (db *userRepository) GetAllUser(perPage int64, offsets int64) (*[]model.UserModel, int64) {
	var userModel []model.UserModel
	var count int64
	errs := db.connection.Table("users").Count(&count).Error
	if errs != nil {
		return nil, 0
	}
	res := db.connection.Table("users").Where("deleted_at IS NULL").Select("id", "name", "email", "address", "phone").Limit(int(perPage)).Offset(int(offsets)).Scan(&userModel)
	if res == nil {
		return nil, 1
	}
	return &userModel, count
}

func (db *userRepository) CreateUser(userModel model.UserModel) (*model.UserModel, error) {
	err := db.connection.Table("users").Create(&userModel).Error
	return &userModel, err
}
