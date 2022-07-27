package login

import (
	"context"
	"github.com/eneskzlcn/dictionary-app-cli/internal/util/http"
)

type RestClient interface {
	Post(ctx context.Context, path string, request, response interface{},headers http.HttpHeaders) error
}

type Client struct {
	restClient RestClient
}

func NewClient(client RestClient) *Client {
	return &Client{restClient: client}
}
func (c *Client) Login(request SignInRequest) SignInResponse {
	return SignInResponse{}
}
