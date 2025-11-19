package enablebankinggo

import "errors"

type (
	// ErrorCode represents error code returned by the API.
	ErrorCode string

	// ErrorResponse represents error response from the API.
	ErrorResponse struct {
		// Message is the error message.
		Message string `json:"message"`

		// Code is the error code, identical to the http response code, if available.
		Code int `json:"code,omitempty"`

		// ErrorCode is the text representation of the error code, if available.
		ErrorCode ErrorCode `json:"error,omitempty"`

		// Detail provides detailed explanation of an error, if available.
		Detail []map[string]any `json:"detail,omitempty"`
	}
)

const (
	// AccessDeniedErrorCode access to this resource is denied. Check services available
	// for your application.
	AccessDeniedErrorCode ErrorCode = "ACCESS_DENIED"

	// AccountDoesNotExistErrorCode no account found matching provided id.
	AccountDoesNotExistErrorCode ErrorCode = "ACCOUNT_DOES_NOT_EXIST"

	// AlreadyAuthorizedErrorCode session is already authorized.
	AlreadyAuthorizedErrorCode ErrorCode = "ALREADY_AUTHORIZED"

	// ASPSPAccountNotAccessibleErrorCode the PSU does not have access to the requested
	// account or it doesn't exist.
	ASPSPAccountNotAccessibleErrorCode ErrorCode = "ASPSP_ACCOUNT_NOT_ACCESSIBLE"

	// ASPSPErrorErrorCode error interacting with ASPSP.
	ASPSPErrorErrorCode ErrorCode = "ASPSP_ERROR"

	// ASPSPPaymentNotAccessibleErrorCode payment can not be requested from the ASPSP.
	ASPSPPaymentNotAccessibleErrorCode ErrorCode = "ASPSP_PAYMENT_NOT_ACCESSIBLE"

	// ASPSPPsuActionRequiredErrorCode PSU action is required to proceed.
	ASPSPPsuActionRequiredErrorCode ErrorCode = "ASPSP_PSU_ACTION_REQUIRED"

	// ASPSPRateLimitExceededErrorCode ASPSP Rate limit exceeded.
	ASPSPRateLimitExceededErrorCode ErrorCode = "ASPSP_RATE_LIMIT_EXCEEDED"

	// ASPSPTimeoutErrorCode timeout interacting with ASPSP.
	ASPSPTimeoutErrorCode ErrorCode = "ASPSP_TIMEOUT"

	// AuthorizationNotProvidedErrorCode authorization header is not provided.
	AuthorizationNotProvidedErrorCode ErrorCode = "AUTHORIZATION_NOT_PROVIDED"

	// ClosedSessionErrorCode session is closed.
	ClosedSessionErrorCode ErrorCode = "CLOSED_SESSION"

	// DateFromInFutureErrorCode date_from can not be in the future.
	DateFromInFutureErrorCode ErrorCode = "DATE_FROM_IN_FUTURE"

	// DateToWithoutDateFromErrorCode date_from must be provided if date_to provided.
	DateToWithoutDateFromErrorCode ErrorCode = "DATE_TO_WITHOUT_DATE_FROM"

	// ExpiredAuthorizationCodeErrorCode authorization code is expired.
	ExpiredAuthorizationCodeErrorCode ErrorCode = "EXPIRED_AUTHORIZATION_CODE"

	// ExpiredSessionErrorCode session is expired.
	ExpiredSessionErrorCode ErrorCode = "EXPIRED_SESSION"

	// InvalidAccountIDErrorCode either iban or other account identification is required.
	InvalidAccountIDErrorCode ErrorCode = "INVALID_ACCOUNT_ID"

	// InvalidHostErrorCode invalid host.
	InvalidHostErrorCode ErrorCode = "INVALID_HOST"

	// InvalidPaymentErrorCode invalid or expired payment provided.
	InvalidPaymentErrorCode ErrorCode = "INVALID_PAYMENT"

	// NoAccountsAddedErrorCode no allowed accounts added to the application.
	NoAccountsAddedErrorCode ErrorCode = "NO_ACCOUNTS_ADDED"

	// PaymentLimitExceededErrorCode the amount value or the the number of transactions
	// exceeds the limit.
	PaymentLimitExceededErrorCode ErrorCode = "PAYMENT_LIMIT_EXCEEDED"

	// PaymentNotFinalizedErrorCode you can not delete a payment that is not finalized
	// or cancelled.
	PaymentNotFinalizedErrorCode ErrorCode = "PAYMENT_NOT_FINALIZED"

	// PaymentNotFoundErrorCode payment not found.
	PaymentNotFoundErrorCode ErrorCode = "PAYMENT_NOT_FOUND"

	// PSUHeaderNotProvidedErrorCode required PSU header not provided.
	PSUHeaderNotProvidedErrorCode ErrorCode = "PSU_HEADER_NOT_PROVIDED"

	// RedirectURINotAllowedErrorCode redirect URI not allowed.
	RedirectURINotAllowedErrorCode ErrorCode = "REDIRECT_URI_NOT_ALLOWED"

	// RevokedSessionErrorCode session is revoked.
	RevokedSessionErrorCode ErrorCode = "REVOKED_SESSION"

	// SessionDoesNotExistErrorCode no session found matching provided id.
	SessionDoesNotExistErrorCode ErrorCode = "SESSION_DOES_NOT_EXIST"

	// TransactionDoesNotExistErrorCode no transaction found matching provided id.
	TransactionDoesNotExistErrorCode ErrorCode = "TRANSACTION_DOES_NOT_EXIST"

	// UnauthorizedAccessErrorCode unauthorized access.
	UnauthorizedAccessErrorCode ErrorCode = "UNAUTHORIZED_ACCESS"

	// UnauthorizedIPErrorCode used IP address is not authorized to access the resource.
	UnauthorizedIPErrorCode ErrorCode = "UNAUTHORIZED_IP"

	// UntrustedPaymentPartyErrorCode either creditor or debtor account is not trusted.
	UntrustedPaymentPartyErrorCode ErrorCode = "UNTRUSTED_PAYMENT_PARTY"

	// WebhookURINotAllowedErrorCode webhook URI not allowed.
	WebhookURINotAllowedErrorCode ErrorCode = "WEBHOOK_URI_NOT_ALLOWED"

	// WrongASPSPProvidedErrorCode wrong ASPSP name provided.
	WrongASPSPProvidedErrorCode ErrorCode = "WRONG_ASPSP_PROVIDED"

	// WrongAuthorizationCodeErrorCode wrong authorization code provided.
	WrongAuthorizationCodeErrorCode ErrorCode = "WRONG_AUTHORIZATION_CODE"

	// WrongContinuationKeyErrorCode wrong continuation key provided.
	WrongContinuationKeyErrorCode ErrorCode = "WRONG_CONTINUATION_KEY"

	// WrongCredentialsProvidedErrorCode wrong credentials provided.
	// nolint:gosec
	WrongCredentialsProvidedErrorCode ErrorCode = "WRONG_CREDENTIALS_PROVIDED"

	// WrongDateIntervalErrorCode date_from should be less than or equal date_to.
	WrongDateIntervalErrorCode ErrorCode = "WRONG_DATE_INTERVAL"

	// WrongRequestParametersErrorCode wrong request parameters provided.
	WrongRequestParametersErrorCode ErrorCode = "WRONG_REQUEST_PARAMETERS"

	// WrongSessionStatusErrorCode wrong session status.
	WrongSessionStatusErrorCode ErrorCode = "WRONG_SESSION_STATUS"

	// WrongTransactionsPeriodErrorCode wrong transactions period requested.
	WrongTransactionsPeriodErrorCode ErrorCode = "WRONG_TRANSACTIONS_PERIOD"
)

func (e ErrorResponse) Error() string {
	return e.Message
}

// IsErrorResponse checks if the provided error is of type [ErrorResponse] and
// returns it along with a boolean indicating the result.
func IsErrorResponse(err error) (*ErrorResponse, bool) {
	var errorResp *ErrorResponse
	if errors.As(err, &errorResp) {
		return errorResp, true
	}

	return nil, false
}
