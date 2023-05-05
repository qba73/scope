package scope

import "unicode"

var oidcScopes = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x21, 0x21, 1},
		{0x23, 0x5B, 1},
		{0x5D, 0x7E, 1},
	},
}

func Valid(s string) bool {
	for _, v := range s {
		if !unicode.Is(oidcScopes, v) {
			return false
		}
	}
	return true
}
