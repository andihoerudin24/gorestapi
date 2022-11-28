package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ApiError struct {
	Field string
	Msg   string
}

func MsgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "failing there is a minimum of characters"
	}
	return ""
}

func BindErrors(err error) interface{} {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), MsgForTag(fe.Tag())}
		}
		return out
	}
	return nil
}
