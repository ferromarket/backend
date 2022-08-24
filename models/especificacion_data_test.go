package models

import "testing"

func TestEspecificacionDataEmptyName(t *testing.T) {
	especificaciondata := EspecificacionData{Valor: ""}
	want := "especificacion data vacio"
	err := especificaciondata.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`EspecificacionNombre.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestEspecificacionDataNotEmptyName(t *testing.T) {
	especificaciondata := EspecificacionData{Valor: "Bueno"}
	var want error = nil
	err := especificaciondata.Validate()
	if err != nil {
		t.Fatalf(`EspecificacionNombre.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
