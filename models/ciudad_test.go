package models

import "testing"

func TestCiudadEmptyName(t *testing.T) {
	ciudad := Ciudad{Nombre: ""}
	want := "nombre vacío"
	err := ciudad.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`ciudad.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestCiudadNotEmptyName(t *testing.T) {
	ciudad := Ciudad{Nombre: "Concepción"}
	var want error = nil
	err := ciudad.Validate()
	if err != nil {
		t.Fatalf(`ciudad.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
