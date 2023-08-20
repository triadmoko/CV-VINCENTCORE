package helpers

import (
	"testing"
)

func Test_validate_CountDigits(t *testing.T) {
	v := &validate{}
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{12345, 5},
		{0, 1},
		{-123, 3},
	}
	for _, tc := range testCases {
		output := v.CountDigits(tc.input)
		if output != tc.expectedOutput {
			t.Errorf("CountDigits(%d) expected %d, but got %d", tc.input, tc.expectedOutput, output)
		}
	}
}

func TestPasswordValidation(t *testing.T) {
    type test struct {
        password string
        expectedError bool
    }

    tests := []test{
        {password: "12345", expectedError: true},
        {password: "ABCDEFGH", expectedError: true},
        {password: "abcdefgh", expectedError: true},
        {password: "aBcdEfgh", expectedError: true},
        {password: "aBcdEfg1", expectedError: false},
        {password: "1234abcd", expectedError: true},
        {password: "1234Abcd", expectedError: false},
        {password: "1234abcdE", expectedError: false},
        {password: "1234AbcdE", expectedError: false},
    }

    for _, testCase := range tests {
        v := validate{}
        err := v.Password(testCase.password)

        if err != nil && !testCase.expectedError {
            t.Errorf("Validation failed for password %s, expected no error but got: %s", testCase.password, err.Error())
        }

        if err == nil && testCase.expectedError {
            t.Errorf("Validation failed for password %s, expected error but got none", testCase.password)
        }
    }
}

