package entities

type Day struct {
	Open   string `json:"open,omitempty" bson:"open,omitempty" validate:"required"`
	Close  string `json:"close,omitempty" bson:"close,omitempty" validate:"required"`
	Closed bool   `json:"closed,omitempty" bson:"closed,omitempty"`
}

type Week struct {
	Content []Day `json:"content,omitempty" bson:"content,omitempty" validate:"required"`
}

func (c *Day) Validate() []ValidateError {
	if c.Closed {
		return nil
	}
	return validate(c)
}
