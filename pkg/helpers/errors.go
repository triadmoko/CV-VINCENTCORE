package helpers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Message map[string]interface{} `json:"message"`
	Success bool                   `json:"success"`
	Code    int                    `json:"code"`
	Data    interface{}            `json:"data"`
}

func NewValidatorError(err error) Error {
	e := Error{
		Success: false,
		Code:    http.StatusBadRequest,
		Data:    nil,
	}
	e.Message = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Message[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}

	return e
}

func AccessForbidden() Error {
	e := Error{}
	e.Message = make(map[string]interface{})
	e.Message["body"] = "access forbidden"
	return e
}

func NotFound() Error {
	e := Error{}
	e.Message = make(map[string]interface{})
	e.Message["body"] = "resource not found"
	return e
}
