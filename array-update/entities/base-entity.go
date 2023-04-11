package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseEntity interface {
	Validate() []ValidateError
	SetId(id primitive.ObjectID)
	GetId() primitive.ObjectID
}
