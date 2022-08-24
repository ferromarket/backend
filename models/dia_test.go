package models

import "testing"

func TestDiaEmptyName(t *testing.T) {
	dia := Dia{Nombre: ""}
	want := "nombre vacío"
	err := dia.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`dia.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestDiaNotEmptyName(t *testing.T) {
	dia := Dia{Nombre: "Lunes"}
	var want error = nil
	err := dia.Validate()
	if err != nil {
		t.Fatalf(`dia.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestDiaTooLong(t *testing.T) {
	dia := Dia{Nombre: "Miércoless"}
	want := "nombre está demasiado largo"
	err := dia.Validate()
	if err == nil || (err != nil && err.Error() != want) {
		t.Fatalf(`dia.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestDiaNotTooLong(t *testing.T) {
	dia := Dia{Nombre: "Miércoles"}
	var want error = nil
	err := dia.Validate()
	if err != want {
		t.Fatalf(`dia.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
