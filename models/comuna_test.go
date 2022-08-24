package models

import "testing"

func TestComunaEmptyName(t *testing.T) {
	comuna := Comuna{Nombre: ""}
	want := "nombre vacío"
	err := comuna.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`comuna.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestComunaNotEmptyName(t *testing.T) {
	comuna := Comuna{Nombre: "Concepción"}
	var want error = nil
	err := comuna.Validate()
	if err != nil {
		t.Fatalf(`comuna.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
