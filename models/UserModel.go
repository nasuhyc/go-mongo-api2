package models

import (
	"time"

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
