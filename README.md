[![Go](https://github.com/qba73/scope/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/scope/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/scope)](https://goreportcard.com/report/github.com/qba73/scope)


# scope

`scope` is a Go library for validating OIDC token scopes. It allows you to verify if tokens meet validation requirements described in the [RFC6749](https://datatracker.ietf.org/doc/html/rfc6749#section-3.3) document.

## Using the Go library

```go
import (
    "fmt"
    "strings"

    "github.com/qba73/scope"
)

func main() {
    tokens := "openid+myscope+email"
    
    for _, token := range strings.Split(tokens, "+") {
        if !scope.Valid(token) {
            fmt.Printf("scope/token %v is not valid\n", token)
        }
        fmt.Printf("scope/token %v is valid\n", token)
    }
}
```

## Bugs and feature requests

If you find a bug in the `scope` library, please open an issue. Similarly, if you'd like a feature added or improved, let me know via an issue.

Pull requests welcome!
