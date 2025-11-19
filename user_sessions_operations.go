package enablebankinggo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type (
	// StartAuthorizationRequest represents request to start user authorization (POST /auth).
	StartAuthorizationRequest struct {
		// Access specifies scope of access to be request from ASPSP and to be confirmed by PSU.
		Access *Access `json:"access"`

		// ASPSP is the ASPSP that PSU is going to be authenticated to.
		ASPSP ASPSP `json:"aspsp"`

		// State is an opaque value used by the client to maintain state between the request and
		// callback. Same string will be returned in query parameter when redirecting to the URL
		// passed via redirect_url parameter
		State string `json:"state"`

		// RedirectURL is the URL that PSU will be redirected to after authorization.
		RedirectURL string `json:"redirect_url"`

		// PSUType is the PSU type which consent is created for
		PSUType PSUType `json:"psu_type,omitempty"`

		// AuthMethod is the desired authorization method (in case ASPSP supports multiple).
		// Supported methods can be obtained from ASPSP auth_methods.
		AuthMethod string `json:"auth_method,omitempty"`

		// Credentials is PSU credentials (User ID, company ID etc.) If not provided, then those are
		// going to be asked from a PSU during authorization.
		Credentials map[string]any `json:"credentials,omitempty"`

		// CredentialsAutoSubmit controls whether user credentials will be autosubmitted (if passed).
		// If set to false then credentials form will be prefilled with passed credentials.
		CredentialsAutoSubmit bool `json:"credentials_autosubmit,omitempty"`

		// Language is the preferred PSU language. Two-letter lowercase language code.
		Language string `json:"language,omitempty"`

		// PSUID is an optional unique identification of a PSU used by the client application. It can
		// be used to match sessions of the same user. Although only hashed value is stored, it is
		// recommended to use anonymised identifiers (i.e. digital ID instead of email or social
		// security number). In case the parameter is not passed by the application, random value will
		// be used.
		PSUID string `json:"psu_id,omitempty"`
	}

	// StartAuthorizationResponse represents response from start user authorization (POST /auth).
	StartAuthorizationResponse struct {
		// URL is the URL to redirect PSU to.
		URL string `json:"url"`

		// AuthorizationID is the PSU authorisation ID, a value used to identify an authorisation session.
		// Please note that another session ID will used to fetch account data.
		AuthorizationID string `json:"authorization_id"`

		// PSUIDHash is the Hashed unique identification of a PSU used by the client application. In case
		// PSU ID is not passed by the client application, the hash is calculated based on a random value.
		// The hash also inherits the application ID, so different hashes will be calculated when using the
		// same PSU ID with different applications.
		PSUIDHash string `json:"psu_id_hash"`
	}

	// AuthorizeSessionRequest represents request to authorize a user session (POST /sessions).
	AuthorizeSessionRequest struct {
		// Code is the authorization code returned when redirecting PSU.
		Code string `json:"code"`
	}

	// AuthorizeSessionResponse represents response from authorizing a user session (POST /sessions).
	AuthorizeSessionResponse struct {
		// SessionID is the ID of the PSU session.
		SessionID string `json:"session_id"`

		// Accounts is the list of authorized accounts.
		Accounts []*AccountResource `json:"accounts"`

		// ASPSP is the ASPSP used with the session.
		ASPSP *ASPSP `json:"aspsp"`

		// PSUType is the PSU type used with the session.
		PSUType PSUType `json:"psu_type"`

		// Access is the scope of access requested from ASPSP and confirmed by PSU.
		Access *Access `json:"access"`
	}

	// GetSessionResponse represents response from GET /sessions/{session_id} endpoint.
	GetSessionResponse struct {
		// Status is the current status of the session.
		Status SessionStatus `json:"status"`

		// Accounts is the list of account IDs available in the session.
		Accounts []string `json:"accounts"`

		// AccountsData account data stored in the session.
		AccountsData []*SessionAccount `json:"accounts_data"`

		// ASPSP is the ASPSP used with the session.
		ASPSP *ASPSP `json:"aspsp"`

		// PSUType is the PSU type used with the session.
		PSUType PSUType `json:"psu_type"`

		// PSUIDHash is the hashed unique identification of a PSU used by the client application. In case
		// PSU ID is not passed by the client application, the hash is calculated based on a random value.
		// The hash also inherits the application ID, so different hashes will be calculated when using the
		// same PSU ID with different applications.
		PSUIDHash string `json:"psu_id_hash"`

		// Access is the scope of access requested from ASPSP and confirmed by PSU.
		Access *Access `json:"access"`

		// Created is the session creation time.
		Created time.Time `json:"created"`

		// Authorized is the session authorization time.
		Authorized *time.Time `json:"authorized,omitempty"`

		// Closed is the session expiration time.
		Closed *time.Time `json:"closed,omitempty"`
	}

	// DeleteSessionRequestParams represents request parameters for DELETE /sessions/{session_id} endpoint.
	DeleteSessionRequestParams struct {
		// Headers represents additional headers to include in the request.
		Headers Header
	}

	// SuccessResponse represents a successful response from the API.
	SuccessResponse struct {
		// Message returns "OK" in case of successful request.
		Message string `json:"message,omitempty"`
	}

	// UserSessionsClient client for user sessions API operations.
	UserSessionsClient interface {
		// StartAuthorization start authorization by getting a redirect link and redirecting a PSU to that link.
		StartAuthorization(ctx context.Context, req *StartAuthorizationRequest) (*StartAuthorizationResponse, error)

		// AuthorizeSession authorize user session by provided authorization code.
		AuthorizeSession(ctx context.Context, req *AuthorizeSessionRequest) (*AuthorizeSessionResponse, error)

		// GetSession get session data by session ID.
		GetSession(ctx context.Context, sessionID string) (*GetSessionResponse, error)

		// DeleteSession delete session by session ID. PSU's bank consent will be closed automatically if possible.
		DeleteSession(ctx context.Context, sessionID string, params *DeleteSessionRequestParams) (*SuccessResponse, error)
	}
)

// StartAuthorization start authorization by getting a redirect link and redirecting a PSU to that link.
func (c *APIClient) StartAuthorization(ctx context.Context, req *StartAuthorizationRequest) (*StartAuthorizationResponse, error) {
	if req == nil {
		return nil, errors.New("req cannot be nil")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodPost, "/auth", req)
	if err != nil {
		return nil, err
	}

	var resp StartAuthorizationResponse
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// AuthorizeSession authorize user session by provided authorization code.
func (c *APIClient) AuthorizeSession(ctx context.Context, req *AuthorizeSessionRequest) (*AuthorizeSessionResponse, error) {
	if req == nil {
		return nil, errors.New("req cannot be nil")
	}

	if req.Code == "" {
		return nil, errors.New("req.Code cannot be empty")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodPost, "/sessions", req)
	if err != nil {
		return nil, err
	}

	var resp AuthorizeSessionResponse
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetSession get session data by session ID.
func (c *APIClient) GetSession(ctx context.Context, sessionID string) (*GetSessionResponse, error) {
	if sessionID == "" {
		return nil, errors.New("sessionID cannot be empty")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/sessions/%s", sessionID), nil)
	if err != nil {
		return nil, err
	}

	var resp GetSessionResponse
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteSession delete session by session ID. PSU's bank consent will be closed automatically if possible.
func (c *APIClient) DeleteSession(ctx context.Context, sessionID string, params *DeleteSessionRequestParams) (*SuccessResponse, error) {
	if sessionID == "" {
		return nil, errors.New("sessionID cannot be empty")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/sessions/%s", sessionID), nil)
	if err != nil {
		return nil, err
	}

	if params != nil && params.Headers != nil {
		params.Headers.FillHTTPHeader(reqHTTP.Header)
	}

	var resp SuccessResponse
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
