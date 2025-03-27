package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client) RequestHelper(method, path string, body any) (*http.Response, error) {
	url := c.BaseURL + path
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
	req.Header.Set("Circle-Token", c.AuthToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Error Reading response: %s", err)
		}
		return nil, fmt.Errorf("%s: %s", res.Status, string(body))
	}

	return res, nil
}
