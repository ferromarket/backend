package utils

import (
	"testing"
)

func TestGenerateVerifier(t *testing.T) {
	want := "K"
	msg, err := generateVerifier("8675309")
	if (msg != want) {
		t.Fatalf(`generateVerifier("8675309") = %q, %v, want match for %#q, nil`, msg, err, want)
	} 
}
