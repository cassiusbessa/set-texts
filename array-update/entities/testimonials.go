package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Testimonial struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Fullname    string             `json:"fullname" bson:"fullname" validate:"required"`
	Depointment string             `json:"depointment" bson:"depointment" validate:"required"`
	Displayed   bool               `json:"displayed" bson:"displayed" validate:"required"`
}

func (c *Testimonial) Validate() []ValidateError {
	return validate(c)
}

func (c *Testimonial) SetId(id primitive.ObjectID) {
	c.Id = id
}

func (c *Testimonial) GetId() primitive.ObjectID {
	return c.Id
}
