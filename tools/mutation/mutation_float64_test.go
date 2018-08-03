package mutation

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
)

func TestMutGaussian(t *testing.T) {
	ind1 := base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.Float64Individual)
	ind3 := MutGaussian(ind2, 0.5, 0.5, 0.5)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
	if !ind2.IsEqual(ind3) {
		t.Errorf("ind2 is not equal to ind3: %v %v", ind2, ind3)
	}
}

func TestMutGaussianWithLimitSlice(t *testing.T) {
	mus := []float64{0.2, 0.4, 0.6, 0.8}
	sigma := []float64{0.2, 0.4, 0.6, 0.8}
	ind1 := base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.Float64Individual)
	ind3 := MutGaussian(ind2, mus, sigma, 0.5)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
}

func TestMutPolyNomialBounded(t *testing.T) {
	ind1 := base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.Float64Individual)
	ind3 := MutPolyNomialBounded(ind2, 0.5, 0.2, 0.8, 0.5)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
}

func TestMutPolyNomialBoundedWithLimitSlice(t *testing.T) {
	mus := []float64{0.2, 0.3, 0.4, 0.3}
	sigma := []float64{0.5, 0.6, 0.7, 0.6}
	ind1 := base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.Float64Individual)
	ind3 := MutPolyNomialBounded(ind2, 0.5, mus, sigma, 0.5)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
}

func TestMutShuffleIndexesFloat64(t *testing.T) {
	ind1 := base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.Float64Individual)
	ind3 := MutShuffleIndexesFloat64(ind2, 0.8)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)

}

func TestMutESLogNormal(t *testing.T) {
	ind1 := base.NewFloat64ESIndividual([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, []float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.Float64ESIndividual)
	ind3 := MutESLogNormal(ind2, 1, 0.5)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
}
