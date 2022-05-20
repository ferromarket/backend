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

func TestGenerateVerifier2(t *testing.T) {
	want := "9"
	msg, err := generateVerifier("86753095")
	if (msg != want) {
		t.Fatalf(`generateVerifier("8675309") = %q, %v, want match for %q, nil`, msg, err, want)
	} 
}

func TestGenerateVerifierInvalidString(t *testing.T) {
	want := "invalid RUT"
	_, err := generateVerifier("8675f309")
	if (err.Error() != want) {
		t.Fatalf(`generateVerifier("8675f309") = %q, want %q, nil`, err.Error(), want)
	}
}
