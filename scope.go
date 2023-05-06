// Package
package scope

import (
	"strings"
	"unicode"
)

var oidcScopes = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x21, 0x21, 1},
		{0x23, 0x5B, 1},
		{0x5D, 0x7E, 1},
	},
}

// Valid takes a string representing token scope
// and validates if the token conforms to RFC6749.
// If the token does not contain invalid characters
// it returns true, false otherwise.
//
// Ref. https://datatracker.ietf.org/doc/html/rfc6749#section-3.3
func Valid(s string) bool {
	for _, v := range s {
		if !unicode.Is(oidcScopes, v) {
			return false
		}
	}
	return true
}

// ValidOIDC takes a string representing OIDC scope
// and validates if all tokens in the scope conforms to RFC6749.
// The input string (scope) is expected to have the following formats:
//
// `openid+scope1+scope2`
// `scope1+openid+scope2`
//
// Tokens should be separated by `+` sign. Order of tokens does not matter.
// ValidOIDC checks if the mandatory token `openid` is present
// in the scope. If it is not present the func returns false.
// If any token in the scope contains invalid characters
// the func will return false.
//
// Ref.
//   - https://openid.net/specs/openid-connect-core-1_0.html section 3.1.2.1
//     (Authentication Request)
//   - https://datatracker.ietf.org/doc/html/rfc6749#section-3.3
func ValidOIDC(s string) bool {
	if !strings.Contains(s, "openid") {
		return false
	}
	for _, token := range strings.Split(s, "+") {
		if !Valid(token) {
			return false
		}
	}
	return true
}
