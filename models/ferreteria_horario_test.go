package models

import "testing"

func TestFerreteriaHorarioValidate(t *testing.T) {
	ferreteriaHorario := FerreteriaHorario{}
	var want error = nil
	err := ferreteriaHorario.Validate()
	if err != nil {
		t.Fatalf(`ferreteriaHorario.Validate() = %v, want match for %#q, nil`, err, want)
	}
}
