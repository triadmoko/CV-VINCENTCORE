package userModel

import (
	"net/mail"
	"vincentcore_interview/pkg/helpers"
)

func (req RequestUser) Validate() []helpers.ResponseError {
	var errors []helpers.ResponseError

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		msg := helpers.ResponseError{
			Field: "email",
			Message: []interface{}{
				"invalid email",
			},
		}
		errors = append(errors, msg)
	}

	return errors
}
func (req RequestUserLogin) Validate() []helpers.ResponseError {
	var errors []helpers.ResponseError
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		msg := helpers.ResponseError{
			Field: "email",
			Message: []interface{}{
				"invalid email",
			},
		}
		errors = append(errors, msg)
	}
	if req.Password == "" {
		msg := helpers.ResponseError{
			Field: "password",
			Message: []interface{}{
				"invalid password",
				"password required",
			},
		}
		errors = append(errors, msg)
	}
	return errors
}
