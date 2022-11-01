package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseModel struct {
	Code    int32
	Success bool
	Message string
	Error   error `json:"-"`
}

type Response struct {
	C *gin.Context
}

func NewResponse(ctx *gin.Context) Response {
	return Response{C: ctx}
}

func (res Response) ResponseFormatter(code int, message string, err error, result map[string]interface{}) {
	ctx := res.C

	if code < 400 {
		ctx.AbortWithStatusJSON(code, gin.H{
			"code":    code,
			"success": true,
			"message": message,
			"error":   nil,
			"data":    result,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"code":    code,
		"success": false,
		"message": message,
		"error":   err,
		"data":    result,
	})
}
