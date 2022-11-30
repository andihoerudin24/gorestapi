package validation

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type CreatePostValidation struct {
	Title   string                `json:"title"   form:"title"   binding:"required,min=2,max=30" `
	Content string                `json:"content" form:"content" binding:"required,min=2,max=500"`
	Slug    string                `json:"slug"    form:"slug"    binding:"required,min=2,max=250"`
	Image   *multipart.FileHeader `json:"image"   form:"image"`
	UserId  int                   `json:"user_id" form:"user_id" binding:"required"`
}

func NewCreatePostValidation() *CreatePostValidation {
	return &CreatePostValidation{}
}

func (createpostValidate *CreatePostValidation) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(&createpostValidate)
	if err != nil {
		return err
	}
	return
}
