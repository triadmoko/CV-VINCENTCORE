package helpers

import (
	"errors"
	"regexp"
)

func ValidatePhone(phone string) (bool, error) {
	req, err := regexp.Compile("^(\\+62|62|0)8[1-9][0-9]{6,9}$")
	if err != nil {
		return false, errors.New("Invalid phone number not match")
	}
	if !req.Match([]byte(phone)) {
		return false, errors.New("Invalid phone number not match")
	}
	return true, nil
}
