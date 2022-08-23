package utils

import (
	"regexp"
	"testing"
)

func TestValidJWT(t *testing.T) {
	token, err := GenerateJWT("1", "123456789", "user@ferromarket.cl")
	if err != nil {
		t.Fatalf(`GenerateJWT("1", "123456789", "user@ferromarket.cl") = failed with error %v`, err)
	}

	err = ValidateToken(token)
	if err != nil {
		t.Fatalf(`ValidateToken(token) = failed with error %q`, err)
	}
}

func TestInvalidJWT(t *testing.T) {
	want := "token contains an invalid number of segments"
	token := "holamundo"
	err := ValidateToken(token)
	if err.Error() != want {
		t.Fatalf(`ValidateToken(token) = failed with error %q`, err)
	}
}

func TestExpiredJWT(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lIiwiZW1haWwiOiJ1c2VyQGZlcnJvbWFya2V0LmNsIiwiZXhwIjoxNjYwMDk4NTA0fQ.YoCwoEd_pgzvuirTAcJFZ_yfNk4vJHuswP8td7Lh0yc"
	err := ValidateToken(token)
	matched, _ := regexp.MatchString("(token is expired by .*)", err.Error())
	if !matched {
		t.Fatalf(`ValidateToken(token) = failed with error %q`, err)
	}
}

func TestModifiedJWTPayload(t *testing.T) {
	want := "signature is invalid"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNyb21lciIsImVtYWlsIjoiY2hyaXNAY3JvbWVyLmNsIiwiZXhwIjoxNjYwMDk4NTA0fQ.YoCwoEd_pgzvuirTAcJFZ_yfNk4vJHuswP8td7Lh0yc"
	err := ValidateToken(token)
	if err.Error() != want {
		t.Fatalf(`ValidateToken(token) = failed with error %q`, err)
	}
}

func TestWrongJWTSecret(t *testing.T) {
	want := "signature is invalid"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lIiwiZW1haWwiOiJ1c2VyQGZlcnJvbWFya2V0LmNsIiwiZXhwIjoxNjYwMDk4NTA0fQ.OQfrQA3RTA4LCwKOUA9qRTT2WOJ2xrdo4IWTq47nFzk"
	err := ValidateToken(token)
	if err.Error() != want {
		t.Fatalf(`ValidateToken(token) = failed with error %q`, err)
	}
}

func TestJWTHeader(t *testing.T) {
	token, err := GenerateJWT("1", "123456789", "user@ferromarket.cl")
	if err != nil {
		t.Fatalf(`GenerateJWT("1", "123456789", "user@ferromarket.cl") = failed with error %v`, err)
	}

	matched, _ := regexp.MatchString("(eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9\\..*\\..*)", token)
	if !matched {
		t.Fatalf(`GenerateJWT("1", "123456789", "user@ferromarket.cl") = failed with an invalid header`)
	}
}
