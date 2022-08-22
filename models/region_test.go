package models

import "testing"

func TestRegionEmptyName(t *testing.T) {
	region := Region{Nombre: ""}
	want := "nombre vacío"
	err := region.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`region.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestRegionNotEmptyName(t *testing.T) {
	region := Region{Nombre: "Biobio", Codigo: "VIII"}
	var want error = nil
	err := region.Validate()
	if err != want {
		t.Fatalf(`region.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestRegionEmptyCode(t *testing.T) {
	region := Region{Nombre: "Bíobío"}
	want := "código vacío"
	err := region.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`region.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestRegionNotEmptyCode(t *testing.T) {
	region := Region{Nombre: "Bíobío", Codigo: "VIII"}
	var want error = nil
	err := region.Validate()
	if err != want {
		t.Fatalf(`region.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestRegionTooLong(t *testing.T) {
	region := Region{Nombre: "Bíobío", Codigo: "VIIII"}
	want := "código demasiado grande"
	err := region.Validate()
	if err == nil || (err != nil && err.Error() != want) {
		t.Fatalf(`region.Validate() = %v, want match for %#q, nil`, err, want)
	}
}

func TestRegionNotTooLong(t *testing.T) {
	region := Region{Nombre: "Bíobío", Codigo: "VIII"}
	var want error = nil
	err := region.Validate()
	if err != want {
		t.Fatalf(`region.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
