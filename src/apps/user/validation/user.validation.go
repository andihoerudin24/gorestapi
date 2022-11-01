package validation

import (
	"github.com/gin-gonic/gin"
)

type CreateUserValidation struct {
	Name    string `json:"name"   form:"name" binding:"required"`
	Email   string `json:"email"  form:"email" binding:"required,email"`
	Address string `json:"address" form:"address" binding:"required"`
	Phone   string `json:"phone" form:"phone" binding:"required"`
}

func NewCreateUserValidation() *CreateUserValidation {
	return &CreateUserValidation{}
}

func (userValidation *CreateUserValidation) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindJSON(&userValidation)
	if err != nil {
		return err
	}
	return
}
