package helpers

import (
	"github.com/go-playground/validator/v10"
)

// Response API model
type Response struct {
	// value message success and error message
	Message string `json:"message"`
	// response to code http 200, 404, 500, etc
	Code int `json:"code"`
	// standart status message success and failed
	Success bool `json:"success"`
	// send data format json, if request success.
	// or request data error send message failed data
	Data interface{} `json:"data"`
	// Pagination nullable
}
type ResponseApiPagination struct {
	// value message success and error message
	Message string `json:"message"`
	// response to code http 200, 404, 500, etc
	Code int `json:"code"`
	// standart status message success and failed
	Success bool `json:"success"`
	// send data format json, if request success.
	// or request data error send message failed data
	Data interface{} `json:"data"`
	// Pagination nullable
	Pagination Paginator `json:"pagination,omitempety"`
}

func ResponseApi(message string, code int, success bool, data interface{}) Response {
	// Respon API format json,
	response := Response{
		Message: message,
		Code:    code,
		Success: success,
		Data:    data,
	}
	return response
}
func ResponseApiWithPagination(message string, code int, success bool, data interface{}, pagination Paginator) ResponseApiPagination {
	// Respon API format json,
	response := ResponseApiPagination{
		Message:    message,
		Code:       code,
		Success:    success,
		Data:       data,
		Pagination: pagination,
	}
	return response
}
func FormatterError(err error) []ResponseError {
	var errs []ResponseError
	for _, e := range err.(validator.ValidationErrors) {
		mapping := ResponseError{
			Field: e.Field(),
		}
		mapping.Message = append(mapping.Message, e.ActualTag())
		errs = append(errs, mapping)
	}
	return errs
}

type ResponseError struct {
	Field   string        `json:"field"`
	Message []interface{} `json:"message"`
}
