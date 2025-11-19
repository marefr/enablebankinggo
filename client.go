package enablebankinggo

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	// ClientDefaultAPIBaseURL is the default base URL for the Enable Banking API.
	ClientDefaultAPIBaseURL = "https://api.enablebanking.com"

	// ClientDefaultTokenTTL is the default token time-to-live (TTL) in seconds (1 hour).
	ClientDefaultTokenTTL = 3600

	// ClientMaximumTokenTTL is the maximum token time-to-live (TTL) in seconds (24 hours).
	ClientMaximumTokenTTL = 86400

	// ClientDefaultTokenTTLExtraTime is the extra time added to the token TTL to account for clock skew.
	ClientDefaultTokenTTLExtraTime = 10 * time.Second
)

// Option represents a configuration option for the client.
type Option func(*APIClient)

// WithBaseURL sets a custom base URL for the Enable Banking API client.
func WithBaseURL(baseURL string) Option {
	return func(c *APIClient) {
		c.baseURL = strings.TrimSuffix(baseURL, "/")
	}
}

// WithHTTPClient sets a custom HTTP client for the Enable Banking API client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *APIClient) {
		c.httpClient = httpClient
	}
}

// WithHTTPTransport sets a custom HTTP transport for the client.
func WithHTTPTransport(transport http.RoundTripper) Option {
	return func(c *APIClient) {
		c.httpClient.Transport = transport
	}
}

// WithTokenTTL sets a custom token time-to-live (TTL) in seconds. Default is [ClientDefaultTokenTTL] seconds. Maximum is [ClientMaximumTokenTTL] seconds.
func WithTokenTTL(ttl int) Option {
	if ttl <= 0 || ttl > ClientMaximumTokenTTL {
		panic("token TTL must be between 1 and 86400 seconds")
	}

	return func(c *APIClient) {
		c.authorizer.tokenTTL = int64(ttl)
	}
}

// WithTokenTTLExtraTime sets a custom extra time duration to account for clock skew when validating token expiration. Default is [ClientDefaultTokenTTLExtraTime].
func WithTokenTTLExtraTime(extraTime time.Duration) Option {
	return func(c *APIClient) {
		c.authorizer.extraTTL = extraTime
	}
}

// WithHeaders sets additional headers to include in every request made by the client.
func WithHeaders(headers Header) Option {
	return func(c *APIClient) {
		for k, v := range headers {
			c.headers.Set(k, v)
		}
	}
}

// WithPSUIPAddressHeader sets the [PSUIPAddressHeaderKey] header to include in every request made by the client.
func WithPSUIPAddressHeader(ipAddress string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUIPAddressHeaderKey, ipAddress)
	}
}

// WithPSUUserAgentHeader sets the [PSUUserAgentHeaderKey] header to include in every request made by the client.
func WithPSUUserAgentHeader(userAgent string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUUserAgentHeaderKey, userAgent)
	}
}

// WithPSURefererHeader sets the [PSURefererHeaderKey] header to include in every request made by the client.
func WithPSURefererHeader(referer string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSURefererHeaderKey, referer)
	}
}

// WithPSUAcceptHeader sets the [PSUAcceptHeaderKey] header to include in every request made by the client.
func WithPSUAcceptHeader(accept string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUAcceptHeaderKey, accept)
	}
}

// WithPSUAcceptCharset sets the [PSUAcceptCharsetHeaderKey] header to include in every request made by the client.
func WithPSUAcceptCharset(acceptCharset string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUAcceptCharsetHeaderKey, acceptCharset)
	}
}

// WithPSUAcceptEncoding sets the [PSUAcceptEncodingHeaderKey] header to include in every request made by the client.
func WithPSUAcceptEncoding(acceptEncoding string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUAcceptEncodingHeaderKey, acceptEncoding)
	}
}

// WithPSUAcceptLanguage sets the [PSUAcceptLanguageHeaderKey] header to include in every request made by the client.
func WithPSUAcceptLanguage(acceptLanguage string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUAcceptLanguageHeaderKey, acceptLanguage)
	}
}

// WithPSUGeoLocationHeader sets the [PSUGeoLocationHeaderKey] header to include in every request made by the client.
func WithPSUGeoLocationHeader(geoLocation string) Option {
	return func(c *APIClient) {
		c.headers.Set(PSUGeoLocationHeaderKey, geoLocation)
	}
}

// NewClientWithKeyFile creates a new Enable Banking API client with the provided application ID, private key file path, and options.
// If no options are provided, the client will use default settings of [ClientDefaultAPIBaseURL], [ClientDefaultTokenTTL], and [ClientDefaultTokenTTLExtraTime].
func NewClientWithKeyFile(applicationID, privateKeyPath string, options ...Option) (*APIClient, error) {
	privateKey, err := loadPrivateKeyFromFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key from file: %w", err)
	}

	return NewClient(applicationID, privateKey, options...)
}

// NewClient creates a new Enable Banking API client with the provided application ID, private key, and options.
// If no options are provided, the client will use default settings of [ClientDefaultAPIBaseURL], [ClientDefaultTokenTTL], and [ClientDefaultTokenTTLExtraTime].
func NewClient(applicationID string, privateKey *rsa.PrivateKey, options ...Option) (*APIClient, error) {
	if applicationID == "" {
		return nil, errors.New("application ID cannot be empty")
	}

	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}

	c := &APIClient{
		baseURL:    ClientDefaultAPIBaseURL,
		httpClient: http.DefaultClient,
		headers:    NewHeaders(),
		authorizer: newAuthorizer(applicationID, privateKey, ClientDefaultTokenTTL, ClientDefaultTokenTTLExtraTime),
	}

	c.httpClient.Timeout = 30 * time.Second

	for _, option := range options {
		option(c)
	}

	return c, nil
}

type APIClient struct {
	baseURL    string
	httpClient *http.Client
	headers    Header
	authorizer *authorizer
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

	c.headers.FillHTTPHeader(req.Header)

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	err = c.authorizer.AuthorizeRequest(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *APIClient) sendRequest(req *http.Request, resp any) error {
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
