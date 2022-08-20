package models

import "testing"

func TestCategoriaEmptyName(t *testing.T) {
	categoria := Categoria{Nombre: ""}
	want := "categoria nombre vacia"
	err := categoria.Validate()
	if err == nil || err.Error() != want {
		t.Fatalf(`categoria.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
