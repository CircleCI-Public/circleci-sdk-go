package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"

	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
)

type Client struct {
	baseURL   string
	client    *retryablehttp.Client
	authToken string
	userAgent string
}

type Options struct {
	userAgent string
}

func NewClientWithOptions(baseURL, authToken string, options *Options) *Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	return &Client{
		baseURL:   baseURL,
		client:    retryClient,
		authToken: authToken,
		userAgent: options.userAgent,
	}
}

func NewClient(baseURL, authToken string) *Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	return &Client{
		baseURL:   baseURL,
		client:    retryClient,
		authToken: authToken,
		userAgent: "circleci-sdk-go",
	}
}

func (c *Client) request(ctx context.Context, url, method string, body, respBody any) (_ *Response, err error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}
	req, err := retryablehttp.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Circle-Token", c.authToken)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer closer.ErrorHandler(res.Body, &err)
	defer func() {
		// This helps with connection pooling (makes sure there's nothing trailing in the HTTP request like newlines)
		_, _ = io.Copy(io.Discard, res.Body)
	}()

	if res.StatusCode >= 400 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error Reading response: %w", err)
		}
		return nil, fmt.Errorf("%s: %s", res.Status, string(body))
	}

	if respBody != nil {
		defer res.Body.Close()
		b, _ := io.ReadAll(res.Body) // Body is now empty

		// Pass the byte slice to Unmarshal instead of the Reader to Decoder
		if err := json.Unmarshal(b, respBody); err != nil {
			return nil, errors.Wrapf(err, "error decoding response body: %s", string(b))
		}
	}

	return &Response{
		StatusCode: res.StatusCode,
	}, nil
}

type Response struct {
	StatusCode int
}

// RequestHelperAbsolute is the same as RequestHelper but allows to do a request to other APIs
func (c *Client) RequestHelperAbsolute(ctx context.Context, method, path string, body, respBody any) (*Response, error) {
	return c.request(ctx, path, method, body, respBody)
}

func (c *Client) RequestHelper(ctx context.Context, method, path string, reqBody, respBody any) (*Response, error) {
	return c.request(ctx, c.baseURL+path, method, reqBody, respBody)
}
