package models

import (
	"testing"
)

func TestFerreteriaEmptyName(t *testing.T) {
	ferreteria := Ferreteria{Nombre: ""}
	want := "nombre vacío"
	err := ferreteria.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`ferreteria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestFerreteriaNotEmptyName(t *testing.T) {
	ferreteria := Ferreteria{Nombre: "My Hardware Store", Direccion: "Whatever", Descripcion: "Whatever"}
	var want error = nil
	err := ferreteria.Validate()
	if err != nil {
		t.Fatalf(`ferreteria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestFerreteriaEmptyDirection(t *testing.T) {
	ferreteria := Ferreteria{Nombre: "Whatever", Direccion: ""}
	want := "dirección vacío"
	err := ferreteria.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`ferreteria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestFerreteriaNotEmptyDirection(t *testing.T) {
	ferreteria := Ferreteria{Nombre: "Whatever", Direccion: "Some direction", Descripcion: "Whatever"}
	var want error = nil
	err := ferreteria.Validate()
	if err != nil {
		t.Fatalf(`ferreteria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}


func TestFerreteriaEmptyDescription(t *testing.T) {
	ferreteria := Ferreteria{Nombre: "Whatever", Direccion: "Whatever", Descripcion: ""}
	want := "descripción vacío"
	err := ferreteria.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`ferreteria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestFerreteriaNotEmptyDescription(t *testing.T) {
	ferreteria := Ferreteria{Nombre: "Whatever", Direccion: "Whatever", Descripcion: "Hardware store description"}
	var want error = nil
	err := ferreteria.Validate()
	if err != nil {
		t.Fatalf(`ferreteria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
