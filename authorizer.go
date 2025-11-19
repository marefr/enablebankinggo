package enablebankinggo

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type authorizer struct {
	applicationID string
	privateKey    *rsa.PrivateKey
	tokenTTL      int64
	extraTTL      time.Duration
	m             sync.RWMutex
	token         string
	expiresAt     time.Time
}

func newAuthorizer(applicationID string, privateKey *rsa.PrivateKey, tokenTTL int, extraTTL time.Duration) *authorizer {
	return &authorizer{
		applicationID: applicationID,
		privateKey:    privateKey,
		tokenTTL:      int64(tokenTTL),
		extraTTL:      extraTTL,
	}
}

func (a *authorizer) AuthorizeRequest(req *http.Request) error {
	a.m.RLock()
	if a.token != "" && time.Now().Add(a.extraTTL).Before(a.expiresAt) {
		token := a.token
		a.m.RUnlock()
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
	a.m.RUnlock()

	a.m.Lock()
	defer a.m.Unlock()

	if a.token != "" && time.Now().Add(a.extraTTL).Before(a.expiresAt) {
		req.Header.Set("Authorization", "Bearer "+a.token)
		return nil
	}

	err := a.generateJWT()
	if err != nil {
		return fmt.Errorf("failed to create JWT: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+a.token)
	return nil
}

func (a *authorizer) generateJWT() error {
	header, err := getJwtHeader(a.applicationID)
	if err != nil {
		return err
	}
	body, expiresAt, err := getJwtBody(a.tokenTTL)
	if err != nil {
		return err
	}
	signBody := fmt.Sprintf("%s.%s", header, body)
	signature, err := sign(a.privateKey, []byte(signBody))
	if err != nil {
		return err
	}

	a.token = fmt.Sprintf("%s.%s", signBody, signature)
	a.expiresAt = expiresAt
	return nil
}
