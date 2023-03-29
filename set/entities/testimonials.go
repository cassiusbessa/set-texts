package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Testimonial struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname    string             `json:"fullname" bson:"fullname" validate:"required"`
	Depointment string             `json:"depointment" bson:"depointment" validate:"required"`
	Displayed   bool               `json:"displayed" bson:"displayed"`
}

func (c *Testimonial) Validate() []ValidateError {
	return validate(c)
}
