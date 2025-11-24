# Enable Banking Go Library

A Go library for the Enable Banking API's.

[![License](https://img.shields.io/github/license/marefr/enablebankinggo)](LICENSE)
[![Go.dev](https://pkg.go.dev/badge/github.com/marefr/enablebankinggo)](https://pkg.go.dev/github.com/marefr/enablebankinggo?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/marefr/enablebankinggo)](https://goreportcard.com/report/github.com/marefr/enablebankinggo)
[![CI](https://github.com/marefr/enablebankinggo/actions/workflows/go.yml/badge.svg)](https://github.com/marefr/enablebankinggo/actions/workflows/go.yml)

The following Go packages are included:
- enablebankinggo: Provides a library for the Enable Banking API, that supports  authorizing and retrieving account data and transactions.
- enablebankinggo/controlpanel: Provides a library for the Enable Banking Control Panel API, that supports authorizing and managing API applications programmatically.

Note: Operations related to payment initiation service (PIS) and payments are not supported.

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
    if errResp, ok := enablebankinggo.IsErrorResponse(err); ok {
      b, err := json.MarshalIndent(errResp, "", "  ")
      if err != nil {
        log.Fatalf("failed to marshal error response: %v", err)
      }
      _, err = out.Write(b)
      if err != nil {
        log.Fatalf("failed to write error response output: %v", err)
      }
    }
    log.Fatalf("failed to get application: %v", err)
  }

  b, err := json.MarshalIndent(appResp, "", "  ")
  if err != nil {
    log.Fatalf("failed to marshal response: %v", err)
  }
  _, err = out.Write(b)
  if err != nil {
    log.Fatalf("failed to write response output: %v", err)
  }
}
```

## References
- [Enable Banking website](https://enablebanking.com)
- [API Reference](https://enablebanking.com/docs/api/reference/)
- [Control Panel](https://enablebanking.com/docs/api/control-panel/)

## Disclaimer

This project is an independent work and has no affiliation, association, authorization, or endorsement from Enable Banking Oy (https://enablebanking.com/) or any of its subsidiaries.
All trademarks, service marks, and company names mentioned herein are the property of their respective owners.

The Enable Banking Control Panel API is not offically documented, besides https://github.com/enablebanking/enablebanking-cli, and therefore there's no guarantees the `controlpanel` package works/behaves as expected/intended.

## License

[MIT](https://github.com/marefr/enablebankinggo/blob/main/LICENSE)
