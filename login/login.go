package login

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInRequest struct {
	Email Email `json:"email"`
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
		is.Alphanumeric,
	)
}
func(p *Password) String() string {
	return string(*p)
}


const (
	ShortestEmailLength = 3
	LongestEmailLength = 320
	ShortestPasswordLength = 6
	LongestPasswordLength = 21
)
