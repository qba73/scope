package scope_test

import (
	"testing"

	"github.com/qba73/scope"
)

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
