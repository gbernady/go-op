# go-op

The `go-op` package is a simple Go wrapper for the [1Password CLI](https://developer.1password.com/docs/cli/get-started/).

## Status

### Supported Features

- [x] account (list, get)
- [ ] connect
- [ ] document
- [ ] events-api
- [x] groups (list, get)
- [x] item (list, create, get, delete)
- [x] user (list, get, get current, get fingerprint, get public key)
- [x] vault (list, get)
- [x] version

## Installation

```go
import "github.com/gbernady/go-op"
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/gbernady/go-op"
)

func main() {
    cli := &op.CLI{}

    item, err := cli.GetItem("Foo")
    if err != nil {
        panic(err)
    }

    fmt.Println("user", item.Field("username").Value)
    fmt.Println("pass", item.Field("password").Value
}
```

## License

The code is licensed under the [MIT License](./LICENSE).
