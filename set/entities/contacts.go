package entities

type Contacts struct {
	Phone        string `json:"phone" bson:"phone" validate:"required,min=10,max=15"`
	Whatsapp     string `json:"whatsapp" bson:"whatsapp" validate:"required"`
	SocialMedia1 string `json:"socialMedia1" bson:"socialMedia1" validate:"required"`
	SocialMedia2 string `json:"socialMedia2" bson:"socialMedia2" validate:"required"`
}

func (c *Contacts) Validate() []ValidateError {
	return validate(c)
}
