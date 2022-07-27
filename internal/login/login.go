package login

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInRequest struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}
type SignInResponse struct {
	Success bool `json:"success"`
}

type Email string

type Password string

func(e *Email) Validate() error {
	return validation.Validate(string(*e),
		validation.Length(ShortestEmailLength, LongestEmailLength),
		is.EmailFormat,
		)
}

func(e *Email) String() string {
	return string(*e)
}

func(p *Password) Validate() error {
	return validation.Validate(string(*p),
		validation.Length(ShortestPasswordLength, LongestPasswordLength),
		validation.By(func(value interface{}) error {
			 if !IsValidPassword(value.(string)) {
				 return errors.New("not valid password")
			 }
			 return nil
		}),
	)
}


func(p *Password) String() string {
	return string(*p)
}

const (
	UppercasedLetterASCIILowerBound = 65
	UppercasedLetterASCIIUpperBound = 90
	LowercasedLetterASCIILowerBound = 97
	LowercasedLetterASCIIUpperBound = 122
	DecimalCharacterASCIILowerBound = 48
	DecimalCharacterASCIIUpperBound = 57
)

func isDecimalChar(char int32) bool {
	return (DecimalCharacterASCIILowerBound <= char) && (char <= DecimalCharacterASCIIUpperBound)
}
func isUppercaseLetter(char int32) bool {
	return (UppercasedLetterASCIILowerBound <= char) && (char <= UppercasedLetterASCIIUpperBound)
}
func isLowercaseLetter(char int32) bool {
	return (LowercasedLetterASCIILowerBound <= char) && (char <= LowercasedLetterASCIIUpperBound)
}
func IsValidPassword(password string) bool {
	decimalCount := 0
	hasSpecialChar := false
	hasUppercaseLetter := false
	hasLowercaseLetter := false

	for _, char := range password {
		if isDecimalChar(char) {
			decimalCount++
		}else if isLowercaseLetter(char) {
			hasLowercaseLetter = true
		}else if isUppercaseLetter(char) {
			hasUppercaseLetter = true
		}else {
			hasSpecialChar = true
		}
	}
	if decimalCount >= MinDecimalCharInPassword && hasSpecialChar && hasUppercaseLetter && hasLowercaseLetter {
		return true
	}
	return false
}

const (
	ShortestEmailLength = 3
	LongestEmailLength = 320
	ShortestPasswordLength = 6
	LongestPasswordLength = 21
	MinDecimalCharInPassword = 3
)
