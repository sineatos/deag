package base

import (
	"testing"
)

func TestFitness(t *testing.T) {
	fitness1 := NewFitness([]float64{1.0, 2.0})
	fitness2 := NewFitnessWithValues([]float64{1.0, 2.0}, []float64{2.0, 2.0})
	t.Log(fitness1)
	t.Log(fitness1.Clone())
	t.Log(fitness2)
	t.Log(fitness2.Clone())
	if fitness1.Valid() {
		t.Errorf("fitness1 is invalid: %v", fitness1)
	}
	fitness1.SetValues([]float64{1.0, 1.0})
	if fitness1.Greater(fitness2) {
		t.Errorf("fitness1 is less than fitness2: %v\t%v", fitness1, fitness2)
	}
	if fitness1.GreaterEqual(fitness2) {
		t.Errorf("fitness1 is less than fitness2: %v\t%v", fitness1, fitness2)
	}
	if !fitness1.Less(fitness2) {
		t.Errorf("fitness1 is less than fitness2: %v\t%v", fitness1, fitness2)
	}
	if !fitness1.LessEqual(fitness2) {
		t.Errorf("fitness1 is less than fitness2: %v\t%v", fitness1, fitness2)
	}
	if fitness1.Equal(fitness2) {
		t.Errorf("fitness1 isn't equal to fitness2: %v\t%v", fitness1, fitness2)
	}
	if cFitness2 := fitness2.Clone(); !cFitness2.Equal(fitness2) {
		t.Errorf("clone of fitness2 isn't equal to fitness: %v\t%v", fitness2, cFitness2)
	}
	if !fitness2.Dominates(fitness1, nil) {
		t.Errorf("fitness2 is dominates fitness1: %v\t%v", fitness1, fitness2)
	}
	wvalues1, wvalues2 := fitness1.GetWValues(), fitness2.GetWValues()
	t.Log(wvalues1)
	t.Log(wvalues2)
	fitness1.Invalidate()
	if fitness1.Valid() {
		t.Errorf("fitness1 is valid %v", fitness1)
	}
}
