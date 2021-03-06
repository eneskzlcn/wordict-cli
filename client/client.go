package client

import (
	"bytes"
	"context"
	"encoding/json"
	httpUtil "github.com/eneskzlcn/dictionary-app-cli/internal/util/http"
	"io"
	"net/http"
)


type Client struct {
	baseURL string
}
func NewClient(baseUrl string) *Client {
	return &Client{baseURL: baseUrl}
}
func (c *Client) Get(ctx context.Context, path string, response interface{}, headers httpUtil.HttpHeaders) error {
	resp, err := c.sendRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(respBytes, response)
}
func (c *Client) Post(ctx context.Context, path string, request, response interface{}, headers httpUtil.HttpHeaders) error {
	resp, err := c.sendRequest(ctx, http.MethodGet, path, request, headers)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(respBytes, response)
}
func (c *Client) Put(ctx context.Context, path string, request, response interface{}, headers httpUtil.HttpHeaders) error {
	resp, err := c.sendRequest(ctx, http.MethodPut, path, request, headers)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(respBytes, response)
}
func (c *Client) Delete(ctx context.Context, path string, request, response interface{}, headers httpUtil.HttpHeaders) error {
	resp, err := c.sendRequest(ctx, http.MethodDelete, path, request, headers)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(respBytes, response)
}
func (c *Client) sendRequest(ctx context.Context, method, path string, requestBody interface{}, headers httpUtil.HttpHeaders) (*http.Response, error){
	requestBodyBytes,err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx,method,c.baseURL+path, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		request.Header.Add(key,value)
	}
	return http.DefaultClient.Do(request)
}