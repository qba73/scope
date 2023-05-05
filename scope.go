package scope

import "unicode"

var oidcScopes = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x21, 0x21, 1},
		{0x23, 0x5B, 1},
		{0x5D, 0x7E, 1},
	},
}

// Valid takes a string representing OIDC token scope
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
