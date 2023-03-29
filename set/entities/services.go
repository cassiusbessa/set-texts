package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Services struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	MinPrice    float32            `json:"min_price" bson:"min_price" validate:"required"`
}

func (c *Services) Validate() []ValidateError {
	return validate(c)
}
