package zenhub

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.zenhub.io/"
)

// Option is a configuration option for the client.
type Option func(client *Client) error

// Options is a namespace var for configuration options.
var Options = ClientOptions{}

// ClientOptions is a namespace for configuration option methods.
type ClientOptions struct{}

// Secret registers the ZenHub secret.
func (ClientOptions) Secret(secret string) Option {
	return func(c *Client) error {
		c.secret = &secret
		return nil
	}
}

// Client registers a (custom) http client.
func (ClientOptions) Client(client *http.Client) Option {
	return func(c *Client) error {
		c.client = client
		return nil
	}
}

// baseURL registers a custom enterprise url.
func (ClientOptions) BaseURL(baseURL string) Option {
	return func(c *Client) error {
		baseEndpoint, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		if !strings.HasSuffix(baseEndpoint.Path, "/") {
			baseEndpoint.Path += "/"
		}
		c.baseURL = baseEndpoint
		return nil
	}
}

type Client struct {
	client *http.Client
	// Base URL for API requests. Defaults to the public ZenHub API.
	baseURL *url.URL
	secret  *string
}

func NewClient(options ...Option) (*Client, error) {
	c := new(Client)
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, errors.New("error applying option")
		}
	}
	if c.client == nil {
		c.client = &http.Client{}
	}
	if c.baseURL == nil {
		baseURL, _ := url.Parse(defaultBaseURL)
		c.baseURL = baseURL
	}
	return c, nil
}

func (c *Client)  NewRequest(method, relativeURL string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("base url must have a trailing slash, but %q does not", c.baseURL)
	}
	u, err := c.baseURL.Parse(relativeURL)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.secret != nil {
		req.Header.Set("X-Authentication-Token", *c.secret)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)

	if v != nil {
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors
		}
		if decErr != nil {
			err = decErr
		}

	}
	return resp, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return errors.New(fmt.Sprintf("non-200 error returned: %d", r.StatusCode))
}
