package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	MinPrice    float32            `json:"min_price" bson:"min_price" validate:"required"`
}

func (c *Service) Validate() []ValidateError {
	return validate(c)
}

func (c *Service) SetId(id primitive.ObjectID) {
	c.Id = id
}

func (c *Service) GetId() primitive.ObjectID {
	return c.Id
}
