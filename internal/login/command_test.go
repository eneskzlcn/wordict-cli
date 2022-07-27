package login_test

import (
	"github.com/eneskzlcn/dictionary-app-cli/cli"
	"github.com/eneskzlcn/dictionary-app-cli/internal/login"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewLoginCommand(t *testing.T) {
	t.Run("Test given command run func then it should register the func to command when NewLoginCommand called", func(t *testing.T) {
		testData := false
		command := login.NewLoginCommand(func(command *cli.Command, args []string) {
			testData = true
		})
		command.Execute()
		assert.NotNil(t, command)
		assert.Equal(t, true, testData)
	})
	t.Run("Test email and password flags added to the command when NewLoginCommand called", func(t *testing.T) {
		command := login.NewLoginCommand(func(command *cli.Command, args []string) {
		})

		hasEmailFlag := command.HasFlag(login.EmailFlagName)
		hasPasswordFlag := command.HasFlag(login.PasswordFlagName)

		assert.NotNil(t, command)
		assert.Equal(t, true, hasEmailFlag)
		assert.Equal(t, true, hasPasswordFlag)
	})

}