# DataMammoth Go SDK

Official Go client for the [DataMammoth API v2](https://data-mammoth.com/api-docs/reference).

> **Status**: Under development. Not yet published.

## Installation

```bash
go get github.com/datamammoth/dm-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    dm "github.com/datamammoth/dm-go"
)

func main() {
    client := dm.NewClient("dm_your_key_here")
    ctx := context.Background()

    // List active servers
    servers, err := client.Servers.List(ctx, &dm.ServerListParams{
        Status: dm.String("active"),
    })
    if err != nil {
        log.Fatal(err)
    }
    for _, s := range servers.Data {
        fmt.Printf("%s — %s\n", s.Hostname, s.IPAddress)
    }

    // Create a server
    task, err := client.Servers.Create(ctx, &dm.ServerCreateParams{
        ProductID: "prod_abc",
        ImageID:   "img_ubuntu2204",
        Hostname:  "web-01",
    })
    if err != nil {
        log.Fatal(err)
    }
    server, err := task.Wait(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created: %s\n", server.ID)
}
```

## Features

- All 105 API v2 endpoints
- Strongly typed request/response structs
- Context-aware with cancellation support
- Automatic pagination iterators
- Rate limit handling with retry
- API key authentication

## Documentation

- [API Reference](https://data-mammoth.com/api-docs/reference)
- [Getting Started Guide](https://data-mammoth.com/api-docs/guides)
- [Authentication](https://data-mammoth.com/api-docs/guides/authentication)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE).
