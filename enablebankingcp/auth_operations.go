package enablebankingcp

import (
	"context"
	"errors"
	"net/http"
	"net/url"
)

// RelyingpartyGetOOBConfirmationCodeRequest represents the request payload for the RelyingpartyGetOOBConfirmationCode endpoint.
type RelyingpartyGetOOBConfirmationCodeRequest struct {
	RequestType        string `json:"requestType"`
	Email              string `json:"email"`
	ContinueURL        string `json:"continueUrl"`
	CanHandleCodeInApp bool   `json:"canHandleCodeInApp"`
}

// GetOOBConfirmationCodeResponse represents the response from the RelyingpartyGetOOBConfirmationCode endpoint.
type GetOOBConfirmationCodeResponse struct {
	// Email: The email address that the email is sent to.
	Kind string `json:"kind"`
	// Kind: The fixed string "identitytoolkit#GetOobConfirmationCodeResponse".
	Email string `json:"email"`
	// OOBCode: The code to be send to the user.
	OOBCode string `json:"oobCode,omitempty"`
}

// RelyingpartyEmailLinkSigninRequest represents the request payload for the RelyingpartyEmailLinkSignin endpoint.
type RelyingpartyEmailLinkSigninRequest struct {
	// Email: The email address of the user.
	Email string `json:"email,omitempty"`
	// IDToken: Token for linking flow.
	IDToken string `json:"idToken,omitempty"`
	// OOBCode: The confirmation code.
	OOBCode string `json:"oobCode,omitempty"`
}

// EmailLinkSigninResponse represents the response from the RelyingpartyEmailLinkSignin endpoint.
type EmailLinkSigninResponse struct {
	// Email: The user's email.
	Email string `json:"email,omitempty"`
	// ExpiresIn: Expiration time of STS id token in seconds.
	ExpiresIn int64 `json:"expiresIn,omitempty,string"`
	// IDToken: The STS id token to login the newly signed in user.
	IDToken string `json:"idToken,omitempty"`
	// IsNewUser: Whether the user is new.
	IsNewUser bool `json:"isNewUser,omitempty"`
	// Kind: The fixed string "identitytoolkit#EmailLinkSigninResponse".
	Kind string `json:"kind,omitempty"`
	// LocalID: The RP local ID of the user.
	LocalID string `json:"localId,omitempty"`
	// RefreshToken: The refresh token for the signed in user.
	RefreshToken string `json:"refreshToken,omitempty"`
}

// RefreshTokenResponse represents the response from the token refresh endpoint.
type RefreshTokenResponse struct {
	// AccessToken: The access token for the signed in user.
	AccessToken string `json:"access_token"`

	// ExpiresIn: Expiration time of token in seconds.
	ExpiresIn int64 `json:"expires_in,string"`

	// TokenType: The type of the token.
	TokenType string `json:"token_type"`

	// RefreshToken: The refresh token for the signed in user.
	RefreshToken string `json:"refresh_token"`

	// IDToken: The id token for the signed in user.
	IDToken string `json:"id_token"`

	// UserID: The user ID associated with the token.
	UserID string `json:"user_id"`

	// ProjectID: The project ID associated with the token.
	ProjectID string `json:"project_id"`
}

// RelyingpartyGetOOBConfirmationCode initiates the out-of-band confirmation code process.
func (c *APIClient) RelyingpartyGetOOBConfirmationCode(ctx context.Context, req *RelyingpartyGetOOBConfirmationCodeRequest) (*GetOOBConfirmationCodeResponse, error) {
	if req == nil {
		return nil, errors.New("req cannot be nil")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodPost, "/relyingparty/getOobConfirmationCode", req)
	if err != nil {
		return nil, err
	}

	queries := reqHTTP.URL.Query()
	queries.Add("key", "AIzaSyBn8fvjRYQKslskRaO3cblUjmcyl5b9o-c")
	reqHTTP.URL.RawQuery = queries.Encode()

	var resp GetOOBConfirmationCodeResponse
	err = c.sendUnauthenticatedRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// RelyingpartyEmailLinkSignin completes the email link sign-in process.
func (c *APIClient) RelyingpartyEmailLinkSignin(ctx context.Context, req *RelyingpartyEmailLinkSigninRequest) (*EmailLinkSigninResponse, error) {
	if req == nil {
		return nil, errors.New("req cannot be nil")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodPost, "/relyingparty/emailLinkSignin", req)
	if err != nil {
		return nil, err
	}

	var resp EmailLinkSigninResponse
	err = c.sendUnauthenticatedRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// RefreshToken refreshes the ID token using the provided refresh token.
func (c *APIClient) RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error) {
	values := url.Values{}
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", refreshToken)

	reqHTTP, err := c.newFormDataRequest(ctx, http.MethodPost, "/token", values)
	if err != nil {
		return nil, err
	}

	var resp RefreshTokenResponse
	err = c.sendUnauthenticatedRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
