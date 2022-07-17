package login

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/dictionary-app-cli/util"
)

type LoginClient interface {
	Post(ctx context.Context, path string, request, response interface{},headers util.HttpHeaders) error
}
type Service struct {
	loginClient LoginClient
}

func NewService(loginClient LoginClient) *Service {
	return &Service{loginClient: loginClient}
}
func (s *Service) Login(request SignInRequest) error {
	//var response SignInResponse
	//return s.loginClient.Post(context.Background(),"/login",request, &response, util.JsonContentDefaultHeader)
	fmt.Printf("Logging in to the server with email:%s , password: %s",request.Email, request.Password)
	return nil
}