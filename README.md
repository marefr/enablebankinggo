# Enable Banking Go API Client Library

A Go client library for the Enable Banking API. Supports operations related to account information service (AIS) and retrieval of account data and transactions.

[![License](https://img.shields.io/github/license/marefr/enablebankinggo)](LICENSE)
[![Go.dev](https://pkg.go.dev/badge/github.com/marefr/enablebankinggo)](https://pkg.go.dev/github.com/marefr/enablebankinggo?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/marefr/enablebankinggo)](https://goreportcard.com/report/github.com/marefr/enablebankinggo)
[![CI](https://github.com/marefr/enablebankinggo/actions/workflows/go.yml/badge.svg)](https://github.com/marefr/enablebankinggo/actions/workflows/go.yml)

Note: Operations related to payment initiation service (PIS) and payments are currently not supported.

## Example Usage

```go
package main

import (
  "context"

  "github.com/marefr/enablebankinggo"
)

func main() {
  applicationID := "<application id>"
  privateKeyFile := "<private key file path>"
  client, err := enablebankinggo.NewClientWithKeyFile(applicationID, privateKeyFile)
  if err != nil {
    log.Fatalf("failed to create new client: %v", err)
  }

  appResp, err := client.GetApplication(context.Background())
  if err != nil {
    log.Fatalf("failed to get application: %v", err)
  }

  b, err := json.MarshalIndent(appResp, "", "  ")
  if err != nil {
    log.Fatalf("failed to marshal value: %v", err)
  }
  _, err = out.Write(b)
  if err != nil {
    log.Fatalf("failed to write output: %v", err)
  }
}
```

## References
- [Enable Banking website](https://enablebanking.com)
- [API Reference](https://enablebanking.com/docs/api/reference/)

## Disclaimer

This project is an independent work and has no affiliation, association, authorization, or endorsement from Enable Banking Oy (https://enablebanking.com/) or any of its subsidiaries.
All trademarks, service marks, and company names mentioned herein are the property of their respective owners.

## License

[MIT](https://github.com/marefr/enablebankinggo/blob/main/LICENSE)
