package models

import "testing"

func TestHoraEmptyName(t *testing.T) {
	hora := Hora{Hora: ""}
	want := "hora vacío"
	err := hora.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`hora.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestHoraNotEmptyName(t *testing.T) {
	hora := Hora{Hora: "13:00"}
	var want error = nil
	err := hora.Validate()
	if err != want {
		t.Fatalf(`hora.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestHoraTooLong(t *testing.T) {
	hora := Hora{Hora: "13:450"}
	want := "hora está demasiado largo"
	err := hora.Validate()
	if err == nil || (err != nil && err.Error() != want) {
		t.Fatalf(`hora.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestHoraNotTooLong(t *testing.T) {
	hora := Hora{Hora: "13:45"}
	var want error = nil
	err := hora.Validate()
	if err != want {
		t.Fatalf(`hora.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
