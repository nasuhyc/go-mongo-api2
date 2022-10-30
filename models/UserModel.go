package models

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=32"`
	Surname   string             `json:"surname,omitempty" bson:"surname,omitempty" validate:"required,min=3,max=32"`
	Age       int32              `json:"age,omitempty" bson:"age,omitempty" validate:"required,number"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email,min=6,max=32"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=11,max=11"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt time.Time          `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
	Message     string
}

var validate = validator.New()

func ValidateUserStruct(user User) []*ErrorResponse {
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
