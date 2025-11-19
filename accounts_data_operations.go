package enablebankinggo

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type (
	// GetAccountDetailsRequestParams represents the parameters for the GetAccountDetails API request (GET /accounts/{account_id}/details).
	GetAccountDetailsRequestParams struct {
		// Headers represents additional headers to include in the request.
		Headers Header
	}

	// GetAccountBalancesRequestParams represents the parameters for the GetAccountBalances API request (GET /accounts/{account_id}/balances).
	GetAccountBalancesRequestParams struct {
		// Headers represents additional headers to include in the request.
		Headers Header
	}

	// HalBalances represents the response from retrieving account balances (GET /accounts/{account_id}/balances).
	HalBalances struct {
		Balances []*BalanceResource `json:"balances"`
	}

	// GetAccountTransactionsRequestParams represents the parameters for the GetAccountTransactions API request (GET /accounts/{account_id}/transactions).
	GetAccountTransactionsRequestParams struct {
		// DateFromQueryParam is the date to fetch transactions from (including the date, UTC timezone is assumed).
		DateFromQueryParam time.Time

		// DateToQueryParam is the date to fetch transactions to (including the date, UTC timezone is assumed).
		DateToQueryParam time.Time

		// ContinuationKeyQueryParam is the continuation key, allowing iterate over multiple API pages of transactions.
		ContinuationKeyQueryParam string

		// TransactionStatusQueryParam is the transaction status to filter by.
		TransactionStatusQueryParam TransactionStatus

		// StrategyQueryParam is the strategy how transactions are fetched.
		StrategyQueryParam TransactionsFetchStrategy

		// Headers represents additional headers to include in the request.
		Headers Header
	}

	// HalTransactions represents the response from retrieving account transactions (GET /accounts/{account_id}/transactions).
	HalTransactions struct {
		Transactions    []*Transaction `json:"transactions"`
		ContinuationKey string         `json:"continuation_key,omitempty"`
	}

	// GetTransactionDetailsRequestParams represents the parameters for the GetTransactionDetails API request (GET /accounts/{account_id}/transactions/{transaction_id}).
	GetTransactionDetailsRequestParams struct {
		// Headers represents additional headers to include in the request.
		Headers Header
	}

	// AccountsDataClient client for accounts data API operations.
	AccountsDataClient interface {
		// GetAccountDetails retrieves details of a specific account.
		GetAccountDetails(ctx context.Context, accountID string, params *GetAccountDetailsRequestParams) (*AccountResource, error)

		// GetAccountBalances retrieves balances of a specific account.
		GetAccountBalances(ctx context.Context, accountID string, params *GetAccountBalancesRequestParams) (*HalBalances, error)

		// GetAccountTransactions retrieves transactions of a specific account.
		GetAccountTransactions(ctx context.Context, accountID string, params *GetAccountTransactionsRequestParams) (*HalTransactions, error)

		// GetTransactionDetails retrieves details of a specific transaction for a specific account.
		GetTransactionDetails(ctx context.Context, accountID string, transactionID string, params *GetTransactionDetailsRequestParams) (*Transaction, error)
	}
)

// GetAccountDetails retrieves details of a specific account.
func (c *APIClient) GetAccountDetails(ctx context.Context, accountID string, params *GetAccountDetailsRequestParams) (*AccountResource, error) {
	if accountID == "" {
		return nil, errors.New("accountID cannot be empty")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodGet, "/accounts/"+accountID+"/details", nil)
	if err != nil {
		return nil, err
	}

	if params != nil && params.Headers != nil {
		params.Headers.FillHTTPHeader(reqHTTP.Header)
	}

	var resp AccountResource
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetAccountBalances retrieves balances of a specific account.
func (c *APIClient) GetAccountBalances(ctx context.Context, accountID string, params *GetAccountBalancesRequestParams) (*HalBalances, error) {
	if accountID == "" {
		return nil, errors.New("accountID cannot be empty")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodGet, "/accounts/"+accountID+"/balances", nil)
	if err != nil {
		return nil, err
	}

	if params != nil && params.Headers != nil {
		params.Headers.FillHTTPHeader(reqHTTP.Header)
	}

	var resp HalBalances
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetAccountTransactions retrieves transactions of a specific account.
func (c *APIClient) GetAccountTransactions(ctx context.Context, accountID string, params *GetAccountTransactionsRequestParams) (*HalTransactions, error) {
	if accountID == "" {
		return nil, errors.New("accountID cannot be empty")
	}

	url := "/accounts/" + accountID + "/transactions"
	reqHTTP, err := c.newRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	queryParams := reqHTTP.URL.Query()
	if params != nil {
		if !params.DateFromQueryParam.IsZero() {
			queryParams.Add("date_from", params.DateFromQueryParam.Format(time.DateOnly))
		}

		if !params.DateToQueryParam.IsZero() {
			queryParams.Add("date_to", params.DateToQueryParam.Format(time.DateOnly))
		}

		if params.ContinuationKeyQueryParam != "" {
			queryParams.Add("continuation_key", params.ContinuationKeyQueryParam)
		}

		if params.TransactionStatusQueryParam != "" {
			queryParams.Add("transaction_status", string(params.TransactionStatusQueryParam))
		}

		if params.StrategyQueryParam != "" {
			queryParams.Add("strategy", string(params.StrategyQueryParam))
		}
	}

	reqHTTP.URL.RawQuery = queryParams.Encode()

	if params != nil && params.Headers != nil {
		params.Headers.FillHTTPHeader(reqHTTP.Header)
	}

	var resp HalTransactions
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetTransactionDetails retrieves details of a specific transaction for a specific account.
func (c *APIClient) GetTransactionDetails(ctx context.Context, accountID string, transactionID string, params *GetTransactionDetailsRequestParams) (*Transaction, error) {
	if accountID == "" {
		return nil, errors.New("accountID cannot be empty")
	}

	if transactionID == "" {
		return nil, errors.New("transactionID cannot be empty")
	}

	reqHTTP, err := c.newRequest(ctx, http.MethodGet, "/accounts/"+accountID+"/transactions/"+transactionID, nil)
	if err != nil {
		return nil, err
	}

	if params != nil && params.Headers != nil {
		params.Headers.FillHTTPHeader(reqHTTP.Header)
	}

	var resp Transaction
	err = c.sendRequest(reqHTTP, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
