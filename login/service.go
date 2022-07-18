package login

import (
	"fmt"
)

type Service struct {
	loginClient *LoginClient
}

func NewService(loginClient *LoginClient) *Service {
	return &Service{loginClient: loginClient}
}
func (s *Service) Login(request SignInRequest) error {
	//var response SignInResponse
	//return s.loginClient.Post(context.Background(),"/login",request, &response, util.JsonContentDefaultHeader)
	fmt.Printf("Logging in to the server with email:%s , password: %s",request.Email, request.Password)
	return nil
}