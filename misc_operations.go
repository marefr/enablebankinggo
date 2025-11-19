package enablebankinggo

import (
	"context"
	"net/http"
)

type (
	// GetApplicationResponse represents response from GET /applications endpoint.
	GetApplicationResponse struct {
		// Name is the name of the application.
		Name string `json:"name"`

		// Description is the description of the application, if any.
		Description string `json:"description,omitempty"`

		// KID is the application key id.
		KID string `json:"kid"`

		// Environment is the application environment.
		Environment Environment `json:"environment"`

		// RedirectURLs is the list of allowed redirect urls.
		RedirectURLs []string `json:"redirect_urls"`

		// Active indicates whether the application is active.
		Active bool `json:"active"`

		// Countries is the list of supported countries.
		Countries []string `json:"countries"`

		// Services is the list of supported services.
		Services []Service `json:"services"`
	}

	// GetASPSPsRequestParams represents request parameters for GET /aspsps endpoint.
	GetASPSPsRequestParams struct {
		// CountryQueryParam used to display only ASPSPs from specified country.
		CountryQueryParam string

		// PSUTypeQueryParam used to display only ASPSPs supporting specified PSU type.
		PSUTypeQueryParam PSUType

		// ServiceQueryParam used to display only ASPSPs supporting specified service.
		ServiceQueryParam Service
	}

	// GetASPSPsResponse represents response from GET /aspsps endpoint.
	GetASPSPsResponse struct {
		// ASPSPs is a list of available ASPSPs and countries.
		ASPSPs []*ASPSPData `json:"aspsps"`
	}

	// MiscClient client for miscellaneous API operations.
	MiscClient interface {
		// GetApplication get application associated with provided JWT key ID.
		GetApplication(ctx context.Context) (*GetApplicationResponse, error)

		// GetASPSPs retrieves a list of ASPSPs with their meta information based on provided parameters.
		GetASPSPs(ctx context.Context, params *GetASPSPsRequestParams) (*GetASPSPsResponse, error)
	}
)

// GetApplication retrieves application associated with provided JWT key ID.
func (c *APIClient) GetApplication(ctx context.Context) (*GetApplicationResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/application", nil)
	if err != nil {
		return nil, err
	}

	var resp GetApplicationResponse
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetASPSPs retrieves a list of ASPSPs with their meta information based on provided parameters.
func (c *APIClient) GetASPSPs(ctx context.Context, params *GetASPSPsRequestParams) (*GetASPSPsResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/aspsps", nil)
	if err != nil {
		return nil, err
	}

	queryParams := req.URL.Query()

	if params != nil {
		if params.CountryQueryParam != "" {
			queryParams.Add("country", params.CountryQueryParam)
		}
		if params.PSUTypeQueryParam != "" {
			queryParams.Add("psu_type", string(params.PSUTypeQueryParam))
		}
		if params.ServiceQueryParam != "" {
			queryParams.Add("service", string(params.ServiceQueryParam))
		}
	}

	req.URL.RawQuery = queryParams.Encode()

	var resp GetASPSPsResponse
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
