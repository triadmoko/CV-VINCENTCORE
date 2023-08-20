package helpers

import (
	"errors"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	phone1 := "081234567890"  // valid phone number
	phone2 := "08567890"      // invalid phone number
	phone3 := "628123456789"  // valid phone number
	phone4 := "+628123456789" // valid phone number
	phone5 := "031234567"     // invalid phone number

	// Test with valid phone numbers
	isValid, err := ValidatePhone(phone1)
	if !isValid || err != nil {
		t.Errorf("Error validating valid phone number. Got: isValid=%v, err=%v, want: isValid=%v, err=%v", isValid, err, true, nil)
	}
	isValid, err = ValidatePhone(phone3)
	if !isValid || err != nil {
		t.Errorf("Error validating valid phone number. Got: isValid=%v, err=%v, want: isValid=%v, err=%v", isValid, err, true, nil)
	}
	isValid, err = ValidatePhone(phone4)
	if !isValid || err != nil {
		t.Errorf("Error validating valid phone number. Got: isValid=%v, err=%v, want: isValid=%v, err=%v", isValid, err, true, nil)
	}

	// Test with invalid phone numbers
	isValid, err = ValidatePhone(phone2)
	if isValid || err == nil {
		t.Errorf("Error validating invalid phone number. Got: isValid=%v, err=%v, want: isValid=%v, err=%v", isValid, err, false, errors.New("Invalid phone number not match"))
	}
	isValid, err = ValidatePhone(phone5)
	if isValid || err == nil {
		t.Errorf("Error validating invalid phone number. Got: isValid=%v, err=%v, want: isValid=%v, err=%v", isValid, err, false, errors.New("Invalid phone number not match"))
	}
}
