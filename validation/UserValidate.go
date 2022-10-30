package validation

import "github.com/go-playground/validator/v10"

type User struct {
	Name    string `validate:"required,min=3,max=32"`
	Surname string `validate:"required,min=3,max=32"`
	Age     int32  `validate:"required,number"`
	Email   string `validate:"required,min=3,max=32"`
	Phone   string `validate:"required,min=11,max=11"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
