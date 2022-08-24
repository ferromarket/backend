package models

import "testing"

func TestEspecificacionNombreEmptyName(t *testing.T) {
	especificacionnombre := EspecificacionNombre{Nombre: ""}
	want := "especificacion nombre vacio"
	err := especificacionnombre.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`EspecificacionNombre.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestEspecificacionNombreNotEmptyName(t *testing.T) {
	especificacionnombre := EspecificacionNombre{Nombre: "Bueno"}
	var want error = nil
	err := especificacionnombre.Validate()
	if err != nil {
		t.Fatalf(`EspecificacionNombre.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
