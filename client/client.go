package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthToken  string
}

func NewClient(baseURL, authToken string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		AuthToken:  authToken,
	}
}

func requestHelperFunction(url, token, method string, body any, client *http.Client) (res *http.Response, err error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Circle-Token", token)

	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		defer closer.ErrorHandler(res.Body, &err)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error Reading response: %w", err)
		}
		return nil, fmt.Errorf("%s: %s", res.Status, string(body))
	}
	return res, nil
}

// RequestHelperAbsolute is the same as RequestHelper but allows to do a request to other APIs
func (c *Client) RequestHelperAbsolute(method, path string, body any) (*http.Response, error) {
	return requestHelperFunction(path, c.AuthToken, method, body, c.HTTPClient)
}

func (c *Client) RequestHelper(method, path string, body any) (*http.Response, error) {
	return requestHelperFunction(c.BaseURL+path, c.AuthToken, method, body, c.HTTPClient)
}
