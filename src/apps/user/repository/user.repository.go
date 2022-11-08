package repository

import (
	"gorestapi/src/apps/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(perPage int64, offset int64) (*[]model.UserModel, int64)
	CreateUser(userModel model.UserModel) (*model.UserModel, error)
	FindById(id int64) (*model.UserModel, error)
	Update(id int64, userModel model.UserModel) int64
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

	var datares []interface{}

	if errs != nil {
		return nil, 0
	}

	if count <= 5 {
		offsets = 0
	}

	res := db.connection.Table("users").Where("deleted_at IS NULL").Select("id", "name", "email", "address", "phone", "image").Limit(int(perPage)).Offset(int(offsets)).Scan(&userModel)

	if res == nil {
		return nil, 1
	}

	for index := range userModel {
		datares = append(datares, map[string]interface{}{
			"id":   userModel[index].ID,
			"name": userModel[index].Name,
		})
	}

	return &userModel, count
}

func (db *userRepository) CreateUser(userModel model.UserModel) (*model.UserModel, error) {
	err := db.connection.Table("users").Create(&userModel).Error
	return &userModel, err
}

func (db *userRepository) FindById(id int64) (*model.UserModel, error) {
	var data model.UserModel
	result := map[string]interface{}{}
	res := db.connection.Debug().Table("users").Where("id = ?", id).Find(&data)
	result = map[string]interface{}{
		"data": data.ID,
	}
	if result["data"].(uint) == 0 {
		return nil, res.Error
	} else {
		return &data, nil
	}

}

func (db *userRepository) Update(id int64, userModel model.UserModel) int64 {
	res := db.connection.Model(&userModel).Select("id", "name", "email", "address", "phone").Where("id = ? AND deleted_at is null", id).Updates(map[string]interface{}{
		"name":    userModel.Name,
		"email":   userModel.Email,
		"address": userModel.Address,
		"phone":   userModel.Phone,
	})
	return res.RowsAffected
}
