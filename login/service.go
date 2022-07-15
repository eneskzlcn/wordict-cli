package login

import (
	"context"
	"github.com/eneskzlcn/dictionary-app-cli/util"
)

type LoginClient interface {
	Post(ctx context.Context, path string, request, response interface{},headers map[string]string) error
}
type Service struct {
	loginClient LoginClient
}

func NewService(loginClient LoginClient) *Service {
	return &Service{loginClient: loginClient}
}
func (s *Service) Login(request SignInRequest) error {
	var response SignInResponse
	return s.loginClient.Post(context.Background(),"/login",request, &response, util.JsonContentDefaultHeader)
}