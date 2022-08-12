package utils

import (
	"testing"
	"time"
)

func TestAnyTime(t *testing.T) {
	want := true
	anyTime := AnyTime{}
	msg := anyTime.Match(time.Time{})
	if (msg != want) {
		t.Fatalf(`anyTime.Match(time.Time{}) = false, want match for true`)
	} 
}
