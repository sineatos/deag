package mutation

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
)

func TestMutUniformInt(t *testing.T) {
	ind1 := base.NewIntIndividual([]int{rand.Int(), rand.Int(), rand.Int(), rand.Int()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.IntIndividual)
	ind3 := MutUniformInt(ind2, 10, 20, 0.5)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
	if !ind2.IsEqual(ind3) {
		t.Errorf("ind2 is not equal to ind3: %v %v", ind2, ind3)
	}
}

func TestMutUniformIntWithLimitSlice(t *testing.T) {

	low := []int{10, 20, 30, 40}
	up := []int{12, 50, 80, 100}

	ind1 := base.NewIntIndividual([]int{rand.Int(), rand.Int(), rand.Int(), rand.Int()}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.IntIndividual)
	ind3 := MutUniformInt(ind2, low, up, 0.85)
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
	if !ind2.IsEqual(ind3) {
		t.Errorf("ind2 is not equal to ind3: %v %v", ind2, ind3)
	}
}
