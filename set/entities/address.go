package entities

type Address struct {
	State      string `json:"state" bson:"state" validate:"required"`
	City       string `json:"city" bson:"city" validate:"required"`
	Neiborhood string `json:"neiborhood" bson:"neiborhood" validate:"required"`
	Street     string `json:"street" bson:"street" validate:"required"`
	Number     string `json:"number" bson:"number" validate:"required"`
	Complement string `json:"complement" bson:"complement"`
	ZipCode    string `json:"zipCode" bson:"zipCode" validate:"required,min=8,max=8"`
}

func (c *Address) Validate() []ValidateError {
	return validate(c)
}
