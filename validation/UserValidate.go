package validation

import (
	"go-mongo-api2/models"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
	Message     string
}

var validate = validator.New()

func ValidateUserStruct(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Message = msgForTag(err.StructNamespace())
			errors = append(errors, &element)
		}
	}
	return errors
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "User.Name":
		return "Ad alanı en az 3 karakter olmalıdır"
	case "User.Surname":
		return "Soyad alanı en az 2 karakter olmalıdır."
	case "User.Age":
		return "Yaş alanı boş geçilemez"
	case "User.Email":
		return "Geçerli bir email adresi giriniz.."
	case "User.Phone":
		return "Telefon numarası 11 karakterden oluşan geçerli bir numara olmalıdır."
	}
	return ""
}
