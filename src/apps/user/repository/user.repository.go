package repository

import (
	"fmt"
	"gorestapi/src/apps/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(perPage int64, offset int64) (*[]model.UserModel, int64)
	CreateUser(userModel model.UserModel) (*model.UserModel, error)
	FindById(id int64) (*model.UserModel, error)
	Update(id int64, userModel model.UserModel) int64
	Delete(id int64) error
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

	errs := db.connection.Debug().Table("users").Count(&count).Error

	fmt.Println("datacount", count)

	if errs != nil {
		return nil, 0
	}
	offsets = (offsets - 1) * perPage

	res := db.connection.Table("users").Debug().Where("deleted_at IS NULL").Select("id", "name", "email", "address", "phone", "image").Limit(int(perPage)).Offset(int(offsets)).Scan(&userModel)

	if res == nil {
		return nil, 1
	}

	return &userModel, count
}

func (db *userRepository) CreateUser(userModel model.UserModel) (*model.UserModel, error) {
	err := db.connection.Debug().Table("users").Create(&userModel).Error
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
	var image string
	if userModel.Image != "" {
		image = "image"
	}
	res := db.connection.Model(&userModel).Select("id", "name", "email", "address", "phone", image).Where("id = ? AND deleted_at is null", id).Updates(map[string]interface{}{
		"name":    userModel.Name,
		"email":   userModel.Email,
		"address": userModel.Address,
		"phone":   userModel.Phone,
		"image":   userModel.Image,
	})
	return res.RowsAffected
}

func (db *userRepository) Delete(id int64) error {
	usermodel := model.NewUserModel()
	usermodel.ID = int(id)
	if err := db.connection.Debug().Where("id = ?", id).First(&usermodel).Error; err != nil {
		return err
	}
	ress := db.connection.Debug().Delete(&usermodel)
	return ress.Error
}
