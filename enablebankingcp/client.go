package enablebankingcp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	// ClientDefaultAPIBaseURL is the default base URL for the Enable Banking control panel API.
	ClientDefaultAPIBaseURL = "https://enablebanking.com/api"
)

// ClientOption represents an option for configuring the API client.
type ClientOption func(*APIClient)

// Token represents an authentication token.
type Token struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// WithHTTPClient sets a custom HTTP client for the Enable Banking API client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *APIClient) {
		c.httpClient = httpClient
	}
}

// WithHTTPTransport sets a custom HTTP transport for the client.
func WithHTTPTransport(transport http.RoundTripper) ClientOption {
	return func(c *APIClient) {
		c.httpClient.Transport = transport
	}
}

// WithToken configures the client to use existing token.
func WithToken(token *Token) ClientOption {
	return func(c *APIClient) {
		c.token = token
	}
}

// OnTokenRefreshed configures a callback function to be called after the token have been refreshed.
func OnTokenRefreshed(fn func(token *Token)) ClientOption {
	return func(c *APIClient) {
		c.onTokenRefreshed = fn
	}
}

// APIClient is the Enable Banking control panel API client.
type APIClient struct {
	baseURL          string
	httpClient       *http.Client
	token            *Token
	onTokenRefreshed func(token *Token)
	mu               sync.Mutex
}

// NewClient creates a new Enable Banking control panel API client with default settings.
func NewClient(options ...ClientOption) *APIClient {
	client := &APIClient{
		baseURL:    ClientDefaultAPIBaseURL,
		httpClient: http.DefaultClient,
		token:      &Token{},
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (c *APIClient) newRequest(ctx context.Context, method, url string, reqBody any) (*http.Request, error) {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	var body io.Reader
	if reqBody != nil {
		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		body = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+url, body)
	if err != nil {
		return nil, err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *APIClient) newFormDataRequest(ctx context.Context, method, url string, formData url.Values) (*http.Request, error) {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *APIClient) sendUnauthenticatedRequest(req *http.Request, resp any) error {
	return c.sendRequestInternal(req, resp)
}

func (c *APIClient) sendAuthenticatedRequest(req *http.Request, resp any) error {
	req.Header.Set("Authorization", "Bearer "+c.token.IDToken)

	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return fmt.Errorf("failed to read request body: %w", err)
		}

		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	// Fixme: Multiple assignments to req.Body
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	err := c.sendRequestInternal(req, resp)
	if err != nil {
		if errResp, ok := IsErrorResponse(err); ok && errResp.ErrorObj.Message == "Unauthorized" {
			c.mu.Lock()
			defer c.mu.Unlock()
			if c.token == nil {
				return err
			}

			newTokenResp, refreshErr := c.RefreshToken(req.Context(), c.token.RefreshToken)
			if refreshErr != nil {
				return fmt.Errorf("failed to refresh token: %w", refreshErr)
			}

			c.token.IDToken = newTokenResp.IDToken
			c.token.RefreshToken = newTokenResp.RefreshToken
			c.token.ExpiresIn = newTokenResp.ExpiresIn

			if c.onTokenRefreshed != nil {
				c.onTokenRefreshed(c.token)
			}

			clonedReq := req.Clone(req.Context())
			clonedReq.Header.Set("Authorization", "Bearer "+newTokenResp.IDToken)
			clonedReq.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			return c.sendRequestInternal(clonedReq, resp)
		}

		return err
	}

	return nil
}

func (c *APIClient) sendRequestInternal(req *http.Request, resp any) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 500 {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	if response.StatusCode != 200 {
		var errResp ErrorResponse
		err = json.NewDecoder(response.Body).Decode(&errResp)
		if err != nil {
			return fmt.Errorf("unexpected API error: status code %d", response.StatusCode)
		}

		return &errResp
	}

	if resp != nil {
		return json.NewDecoder(response.Body).Decode(resp)
	}

	return nil
}
