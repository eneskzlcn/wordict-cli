package login

import (
	"context"
	"github.com/eneskzlcn/dictionary-app-cli/util"
)

type Client interface {
	Post(ctx context.Context, path string, request, response interface{},headers util.HttpHeaders) error
}

type LoginClient struct {
	client Client
}

func NewLoginClient(client Client) *LoginClient {
	return &LoginClient{client: client}
}
func (c *LoginClient) Login(request SignInRequest) SignInResponse{
	return SignInResponse{}
}
