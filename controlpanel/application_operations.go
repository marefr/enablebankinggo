package controlpanel

import (
	"context"
	"net/http"
	"net/url"

	"github.com/marefr/enablebankinggo"
)

// RegisterApplicationRequest represents the request payload for registering a new application.
type RegisterApplicationRequest struct {
	// Environment is the application environment.
	Environment enablebankinggo.Environment `json:"environment"`

	// Name is the name of the application.
	Name string `json:"name"`

	// RedirectUrls is the list of allowed redirect URLs.
	RedirectUrls []string `json:"redirect_urls"`

	// Description is the description of the application.
	Description string `json:"description"`

	// PrivacyURL is the URL to the privacy policy of the application.
	PrivacyURL string `json:"privacy_url"`

	// TermsURL is the URL to the terms and conditions of the application.
	TermsURL string `json:"terms_url"`

	// GDPREmail is the data protection email of the application.
	GDPREmail string `json:"gdpr_email"`

	// Certificate is the public key certificate content associated with the application.
	CertificateContent string `json:"certificate"`
}

// LinkApplicationAccountRequest represents the request payload for linking an application account.
type LinkApplicationAccountRequest struct {
	Country string `json:"country"`
	Aspsp   string `json:"aspsp"`
	AppID   string `json:"appId"`
	PsuType string `json:"psuType"`
}

// LinkApplicationAccountResponse represents the response from linking an application account.
type LinkApplicationAccountResponse struct {
	URL             string `json:"url"`
	AuthorizationID string `json:"authorization_id"`
	PsuIDHash       string `json:"psu_id_hash"`
}

// ListApplications retrieves the list of applications.
func (c *APIClient) ListApplications(ctx context.Context) ([]*Application, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/applications", nil)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	err = c.sendAuthenticatedRequest(req, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

// GetApplication get an application by ID.
func (c *APIClient) GetApplication(ctx context.Context, applicationID string) (*Application, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/application/"+applicationID, nil)
	if err != nil {
		return nil, err
	}

	var app *Application
	err = c.sendAuthenticatedRequest(req, &app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// RegisterApplication registers a new application.
func (c *APIClient) RegisterApplication(ctx context.Context, req *RegisterApplicationRequest) (string, error) {
	httpReq, err := c.newRequest(ctx, http.MethodPost, "/applications", req)
	if err != nil {
		return "", err
	}

	var app Application
	err = c.sendAuthenticatedRequest(httpReq, &app)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (c *APIClient) DeleteApplication(ctx context.Context, applicationID string) error {
	req := struct {
		ApplicationID string `json:"appId"`
	}{
		ApplicationID: applicationID,
	}
	httpReq, err := c.newRequest(ctx, http.MethodDelete, "/applications/", &req)
	if err != nil {
		return err
	}

	return c.sendAuthenticatedRequest(httpReq, nil)
}

// LinkApplicationAccount links (whitelists) an account for production tests.
func (c *APIClient) LinkApplicationAccount(ctx context.Context, req *LinkApplicationAccountRequest) (*LinkApplicationAccountResponse, error) {
	data := url.Values{}
	data.Set("country", req.Country)
	data.Set("aspsp", req.Aspsp)
	data.Set("appId", req.AppID)
	data.Set("psuType", req.PsuType)
	data.Set("redirectUrl", "https://enablebanking.com/api/auth_redirect")

	httpReq, err := c.newFormDataRequest(ctx, http.MethodPost, "/link_accounts", data)
	if err != nil {
		return nil, err
	}

	var resp LinkApplicationAccountResponse
	err = c.sendAuthenticatedRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlinkApplicationAccount unlinks (removes from whitelist) an account for production tests.
func (c *APIClient) UnlinkApplicationAccount(ctx context.Context, applicationID string, identificationHash string) error {
	data := url.Values{}
	data.Set("appId", applicationID)
	data.Set("identificationHash", identificationHash)

	httpReq, err := c.newFormDataRequest(ctx, http.MethodPost, "/unlink_accounts", data)
	if err != nil {
		return err
	}

	return c.sendAuthenticatedRequest(httpReq, nil)
}
