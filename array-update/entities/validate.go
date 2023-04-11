package entities

import "github.com/go-playground/validator/v10"

type ValidateError struct {
	Param string `json:"param"`
	Msg   string `json:"msg"`
}

func validate(s BaseEntity) []ValidateError {
	valid := validator.New()
	err := valid.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		out := make([]ValidateError, len(validationErrors))
		for i, e := range validationErrors {
			out[i] = ValidateError{e.Field(), message(e)}
		}
		return out
	}
	return nil
}

func message(e validator.FieldError) string {
	message := map[string]string{
		"lte":      e.Field() + " must be less than or equal to " + e.Param(),
		"gte":      e.Field() + " must be greater than or equal to " + e.Param(),
		"required": e.Field() + " is required",
	}
	return message[e.Tag()]
}
