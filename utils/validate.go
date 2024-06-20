package utils

import (
	// "fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateIncoming(data interface{}) (string, error) {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok { 
			return "Fill all the field", err
		}
		return "Fill all the field", err
	}
	return "", nil
}

