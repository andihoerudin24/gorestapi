package validation

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type CreateUserValidation struct {
	Name    string                `json:"name"    form:"name"    binding:"required"`
	Email   string                `json:"email"   form:"email"   binding:"required,email"`
	Address string                `json:"address" form:"address" binding:"required"`
	Phone   string                `json:"phone"   form:"phone"   binding:"required"`
	Image   *multipart.FileHeader `json:"image"   form:"image"`
}

func NewCreateUserValidation() *CreateUserValidation {
	return &CreateUserValidation{}
}

func (userValidation *CreateUserValidation) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(&userValidation)
	if err != nil {
		return err
	}
	return
}

type UpdateUserValidation struct {
	Name    string                `json:"name"     form:"name" binding:"required"`
	Email   string                `json:"email"    form:"email" binding:"required,email"`
	Address string                `json:"address"  form:"address" binding:"required"`
	Phone   string                `json:"phone"    form:"phone" binding:"required"`
	Image   *multipart.FileHeader `json:"image"    form:"image"`
}

func NewUpdateUserValidation() *UpdateUserValidation {
	return &UpdateUserValidation{}
}

func (updateUserValidation *UpdateUserValidation) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(&updateUserValidation)
	if err != nil {
		return err
	}
	return
}
