package helpers

import (
	"errors"
	"net/mail"
	"strconv"
	"strings"
	"unicode"
)

type Validate interface {
	CountDigits(i int) (count int)
	Password(password string) error
}

type validate struct{}

func NewValidate() *validate {
	return &validate{}
}

func (v *validate) CountDigits(i int) int {
	return len(strings.ReplaceAll(strconv.Itoa(i), "-", ""))
}

// Function to validate password according to specific criteria func
func (v *validate) Password(password string) error {
    var (
        upp, low, num bool // values to track presence of specific password criteria 
    )

    if len(password) < 8 {
        return errors.New("Password must have at least 8 characters and contain at least one uppercase letter, one lowercase letter and one digit")
    }

    for _, c := range password {
        if unicode.IsUpper(c) {
            upp = true
        } else if unicode.IsLower(c) {
            low = true
        } else if unicode.IsNumber(c) {
            num = true
        } else {
            return errors.New("Password must contain only letters and digits")
        }
    }

    if !upp || !low || !num {
        return errors.New("Password must contain at least one uppercase letter, one lowercase letter and one digit")
    }

    return nil
}


func (v *validate) Email(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}
