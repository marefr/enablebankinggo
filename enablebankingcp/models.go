package enablebankingcp

import (
	"time"

	"github.com/marefr/enablebankinggo"
)

// Application represents an application in the Enable Banking control panel.
type Application struct {
	// Name is the name of the application.
	Name string `json:"name"`

	// PrivacyURL is the privacy URL of the application.
	PrivacyURL string `json:"privacy_url,omitempty"`

	// TermsURL is the terms URL of the application.
	TermsURL string `json:"terms_url,omitempty"`

	// GDPREmail is the GDPR contact email of the application.
	GDPREmail string `json:"gdpr_email,omitempty"`

	// Description is the description of the application.
	Description string `json:"description"`

	// KID is the application key id.
	KID string `json:"kid"`

	// Certificate is the certificate associated with the application.
	Certificate *Certificate `json:"certificate"`

	// Environment is the application environment.
	Environment enablebankinggo.Environment `json:"environment"`

	// RedirectUrls is the list of allowed redirect urls.
	RedirectUrls []string `json:"redirect_urls"`

	// Services is the list of services associated with the application.
	Services []enablebankinggo.Service `json:"services"`

	// Active indicates whether the application is active.
	Active bool `json:"active"`

	// WhiteListedAccounts is the list of whitelisted accounts for the application.
	WhiteListedAccounts []*WhiteListedAccount `json:"whitelisted_accounts,omitempty"`

	// Created is the timestamp when the application was created.
	Created time.Time `json:"created"`

	// Creator contains information about the creator of the application.
	Creator struct {
		// Email is the email of the creator.
		Email string `json:"email"`
	} `json:"creator"`

	ConsentOrigins    []map[string]any `json:"consent_origins"`
	WebhookOrigins    []map[string]any `json:"webhook_origins"`
	PendingActionType any              `json:"pendingActionType"`
}

// Certificate represents a certificate associated with an application.
type Certificate struct {
	Source map[string]any `json:"source"`
	JWK    map[string]any `json:"jwk"`
}

// WhiteListedAccount represents a whitelisted account for an application.
type WhiteListedAccount struct {
	// Title is the title of the whitelisted account.
	Title string `json:"title"`

	// IdentificationHash is the identification hash of the whitelisted account.
	IdentificationHash string `json:"identification_hash"`

	// ASPSP is the ASPSP associated with the whitelisted account.
	ASPSP enablebankinggo.ASPSP `json:"aspsp"`

	// Linker is the linker associated with the whitelisted account.
	Linker string `json:"linker"`

	// Created is the timestamp when the whitelisted account was created.
	Created time.Time `json:"created"`
}
