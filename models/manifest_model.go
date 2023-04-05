package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manifest struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Image    string             `json:"image,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Date     string             `json:"date,omitempty" validate:"required"`
}
