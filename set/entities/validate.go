package entities

import "github.com/go-playground/validator/v10"

type ValidateError struct {
	Param string `json:"param"`
	Msg   string `json:"msg"`
}

func validate(s interface{}) []ValidateError {
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
	switch e.Tag() {
	case "lte":
		return e.Field() + " must be less than or equal to " + e.Param()
	case "gte":
		return e.Field() + " must be greater than or equal to " + e.Param()
	case "required":
		return e.Field() + " is required"
	default:
		return ""
	}
}
