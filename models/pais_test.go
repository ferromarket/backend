package models

import "testing"

func TestPaisEmptyName(t *testing.T) {
	pais := Pais{Nombre: ""}
	want := "nombre vacío"
	err := pais.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`pais.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestPaisNotEmptyName(t *testing.T) {
	pais := Pais{Nombre: "Chile", Codigo: "CL"}
	var want error = nil
	err := pais.Validate()
	if err != want {
		t.Fatalf(`pais.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestPaisEmptyCode(t *testing.T) {
	pais := Pais{Nombre: "Chile"}
	want := "código vacío"
	err := pais.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`pais.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestPaisNotEmptyCode(t *testing.T) {
	pais := Pais{Nombre: "Chile", Codigo: "CL"}
	var want error = nil
	err := pais.Validate()
	if err != want {
		t.Fatalf(`pais.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestPaisTooLong(t *testing.T) {
	pais := Pais{Nombre: "Chile", Codigo: "CLP"}
	want := "código demasiado grande"
	err := pais.Validate()
	if err == nil || (err != nil && err.Error() != want) {
		t.Fatalf(`pais.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestPaisNotTooLong(t *testing.T) {
	pais := Pais{Nombre: "Chile", Codigo: "CL"}
	var want error = nil
	err := pais.Validate()
	if err != want {
		t.Fatalf(`pais.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
