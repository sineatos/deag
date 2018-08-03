package base

import (
	"testing"
)

func TestIntIndividual(t *testing.T) {
	ind1 := NewIntIndividual([]int{1, 2, 2}, NewFitness([]float64{-1.0}))
	ind2 := NewIntIndividual([]int{10, 2, 80}, NewFitness([]float64{-1.0}))
	ind3 := NewIntIndividual([]int{1, 2, 2}, NewFitness([]float64{-1.0}))
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

func TestIntESIndividual(t *testing.T) {
	ind1 := NewIntESIndividual([]int{1, 2, 2}, []float64{0.2, 0.6, 0.2}, NewFitness([]float64{-1.0}))
	ind2 := NewIntESIndividual([]int{10, 2, 80}, []float64{0.1, 0.1, 0.8}, NewFitness([]float64{-1.0}))
	ind3 := NewIntESIndividual([]int{1, 2, 2}, []float64{0.2, 0.5, 0.3}, NewFitness([]float64{-1.0}))
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
