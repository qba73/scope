package scope_test

import (
	"fmt"
	"testing"

	"github.com/qba73/scope"
)

func TestValidOIDC_ReturnsTrueOnValidInput(t *testing.T) {
	t.Parallel()

	validInputs := []string{
		"openid",
		"openid email",
		"customScope openid",
		"myscope openid mysecondscope",
	}
	for _, v := range validInputs {
		if !scope.ValidOIDC(v) {
			t.Error(false)
		}
	}

}

func TestValidOIDC_ReturnsFalseOnInvalidInput(t *testing.T) {
	t.Parallel()

	validInputs := []string{
		"openi",
		"email",
		"customScope",
	}
	for _, v := range validInputs {
		if scope.ValidOIDC(v) {
			t.Error(true)
		}
	}

}

func TestValid_ReturnsTrueOnValidInput(t *testing.T) {
	t.Parallel()

	validInputs := []string{
		"abc",
		"openid",
		"email",
		"customScope",
	}
	for _, v := range validInputs {
		if !scope.Valid(v) {
			t.Error(false)
		}
	}
}

func TestValid_ReturnsFalseOnInvalidInput(t *testing.T) {
	t.Parallel()

	invalidInputs := []string{
		"\x20hello",
		"open\x19id",
		"email\x5c",
		"custom\x22Scope",
		"custom\x7f",
	}
	for _, v := range invalidInputs {
		if scope.Valid(v) {
			t.Error(true)
		}
	}
}

func ExampleValidOIDC_validTokens() {
	fmt.Println(scope.ValidOIDC("openid myscope"))
	// Output:
	// true
}

func ExampleValidOIDC_invalidTokens() {
	fmt.Println(scope.ValidOIDC("openid m\x7fyscope"))
	// Output:
	// false
}

func ExampleValidOIDC_missingRequiredToken() {
	fmt.Println(scope.ValidOIDC("secondScope email"))
	// Output:
	// false
}

func ExampleValid_validTokens() {
	fmt.Println(scope.Valid("myscope"))
	// Output:
	// true
}

func ExampleValid_invalidTokens() {
	fmt.Println(scope.Valid("my\x18scope+second\x7fScope"))
	// Output:
	// false
}
