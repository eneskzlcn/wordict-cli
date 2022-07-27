package ask

import "github.com/eneskzlcn/dictionary-app-cli/client"

type AskClient struct {
	client *client.Client
}

func NewAskClient(client *client.Client) *AskClient {
	return &AskClient{client: client}
}
