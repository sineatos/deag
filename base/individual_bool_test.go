package base

import (
	"testing"
)

func TestBoolIndividual(t *testing.T) {
	ind1 := NewBoolIndividual([]bool{true, true, true}, NewFitness([]float64{-1.0}))
	ind2 := NewBoolIndividual([]bool{false, true, false}, NewFitness([]float64{-1.0}))
	ind3 := NewBoolIndividual([]bool{true, true, true}, NewFitness([]float64{-1.0}))
	ind4 := ind1.Clone().(Individual)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 isn't equal to ind3: %v %v", ind1, ind3)
	}
	if !ind3.IsEqual(ind4) {
		t.Errorf("ind3 isn't equal to ind4: %v %v", ind3, ind4)
	}
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
	t.Log(ind4)
	t.Log(ind1.GetChromosome())
	t.Log(ind2.GetChromosome())
	ind3.SetChromosome(ind2.GetChromosome())
	t.Log(ind3)
}
