package login_test

import (
	"github.com/eneskzlcn/dictionary-app-cli/internal/login"
	mocks "github.com/eneskzlcn/dictionary-app-cli/internal/mocks/login"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenValidLoginServiceThenItShouldReturnLoginHandlerWhenNewLoginHandlerCalled(t *testing.T) {
	controller := gomock.NewController(t)
	mockLoginService := mocks.NewMockLoginService(controller)
	loginHandler :=  login.NewHandler(mockLoginService)
	assert.NotNil(t, loginHandler)
}
func TestLoginHandlerInitThenItShouldAddSubcommandsAndFlagsCorrectlyWhenICallInit(t *testing.T) {

	loginHandler :=  login.Handler{}
	loginHandler.Init()

	hasEmailFlag := loginHandler.GetCommand().HasFlag(login.EmailFlagName)
	hasPasswordFlag := loginHandler.GetCommand().HasFlag(login.PasswordFlagName)

	assert.Equal(t, true, hasEmailFlag)
	assert.Equal(t, true, hasPasswordFlag)
}
func TestGivenValidSignInRequestThenItShouldNotReturnErrorWhenLoginCalled(t *testing.T) {
	controller := gomock.NewController(t)
	mockLoginService := mocks.NewMockLoginService(controller)
	loginHandler :=  login.NewHandler(mockLoginService)

	request := &login.SignInRequest{
		Email:    "asd@gmail.com",
		Password: "123@Asfsa1>",
	}
	mockLoginService.EXPECT().Login(request).Return(nil)

	loginHandler.ExecuteCommand()
	//assert.Nil(t, err)

}