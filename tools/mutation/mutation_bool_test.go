package mutation

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
)

func TestMutFlipBit(t *testing.T) {
	ind1 := base.NewBoolIndividual([]bool{true, true, true, true, true, true, true, true, true, true}, base.NewFitness([]float64{-1.0}))
	ind2 := ind1.Clone().(*base.BoolIndividual)
	ind3 := MutFlipBit(ind2, rand.Float64())
	t.Log(ind1)
	t.Log(ind2)
	t.Log(ind3)
	if !ind2.IsEqual(ind3) {
		t.Errorf("ind2 is not equal to ind3: %v %v", ind2, ind3)
	}
}
