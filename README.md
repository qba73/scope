[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/scope.svg)](https://pkg.go.dev/github.com/qba73/scope)
[![Go](https://github.com/qba73/scope/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/scope/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/scope)](https://goreportcard.com/report/github.com/qba73/scope)

# scope

`scope` is a Go library for validating [OIDC token scopes](https://auth0.com/docs/get-started/apis/scopes/openid-connect-scopes). It allows you to verify if tokens meet validation requirements described in the [RFC6749](https://datatracker.ietf.org/doc/html/rfc6749#section-3.3) document.

## Using the Go library

The following example illustrates how to verify the OIDC scope string. The scope must include the `openid` token.

```go
import (
    "fmt"
    "strings"

    "github.com/qba73/scope"
)

func main() {
    scopes := []string{"openid myscope email", "myscope email"}
    
    for _, s := range scopes {
        if !scope.ValidOIDC(s) {
            fmt.Println("invalid scope")
        }
    }
}
```

The following example illustrates how to verify tokens in the scope. Note that func `Valid()` validates
if tokens do not contain [unsupported characters](https://datatracker.ietf.org/doc/html/rfc6749#section-3.3).

```go
import (
    "fmt"
    "strings"

    "github.com/qba73/scope"
)

func main() {
    tokens := "openid myscope email"
    
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
