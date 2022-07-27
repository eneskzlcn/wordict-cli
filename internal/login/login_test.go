package login_test

import (
	"github.com/eneskzlcn/dictionary-app-cli/internal/login"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenValidPasswordThenItShouldReturnTrueWhenIsValidPasswordCalled(t *testing.T) {
	password := "123A!zsgL"

	success := login.IsValidPassword(password)

	assert.Equal(t, success,true)
}
func TestGivenInvalidPasswordThenItShouldReturnFalseWhenIsValidPasswordCalled(t *testing.T) {
	password := "123"
	success := login.IsValidPassword(password)
	assert.Equal(t, false,success)

	password = "123es!"
	success = login.IsValidPassword(password)
	assert.Equal(t, false,success)

	password = "123eS"
	success = login.IsValidPassword(password)
	assert.Equal(t, false,success)
}
func TestValidatePassword(t *testing.T) {
	t.Run("Given valid password then it should validate password when validate called", func(t *testing.T) {
		password := login.Password("1A34!zSgL,")
		err := password.Validate()
		assert.Nil(t, err)
	})

	t.Run("Given invalid password then it should return error when validate called", func(t *testing.T) {
		password := login.Password("123")
		err := password.Validate()
		assert.NotNil(t, err)

		password = "123es!"
		err = password.Validate()
		assert.NotNil(t, err)

		password = "123eS"
		err = password.Validate()
		assert.NotNil(t, err)
	})

}
