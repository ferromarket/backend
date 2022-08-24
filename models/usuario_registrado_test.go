package models

import (
	"testing"
)

func TestUsuarioEmptyName(t *testing.T) {
	usuario := Usuario{Nombres: ""}
	want := "Nombre vacío"
	err := usuario.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`usuario.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestUsuarioNotEmptyName(t *testing.T) {
	usuario := Usuario{Nombres: "juan Heriberto", Direccion: "Algo"}
	var want error = nil
	err := usuario.Validate()
	if err != nil {
		t.Fatalf(`usuario.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
