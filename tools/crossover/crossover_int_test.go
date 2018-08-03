package crossover

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/tools/inits"
)

const (
	idims int = 10
)

// declared in crossover_float64_test.go
// func generateFitness() *base.Fitness {
// 	return base.NewFitness([]float64{-1})
// }

func genInt() int {
	return rand.Intn(idims)
}

func generateIntIndividual() *base.IntIndividual {
	return base.NewIntIndividual(inits.GenerateIntSliceRepeat(genInt, idims), generateFitness())
}

func generateIntESIndividual() *base.IntESIndividual {
	ind := inits.GenerateIntSliceRepeat(genInt, idims)
	st := inits.GenerateFloat64SliceRepeat(rand.Float64, idims)
	return base.NewIntESIndividual(ind, st, generateFitness())
}

func TestCxOnePointInt(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxOnePointInt(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxTwoPointInt(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxTwoPointInt(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxUniformInt(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxUniformInt(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxPartialyMatched(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxPartialyMatched(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxUniformPartialyMatched(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxUniformPartialyMatched(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxOrdered(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxOrdered(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxMessyOnePointInt(t *testing.T) {
	ind1, ind2 := generateIntIndividual(), generateIntIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxMessyOnePointInt(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxESTwoPointInt(t *testing.T) {
	ind1, ind2 := generateIntESIndividual(), generateIntESIndividual()
	t.Log(ind1)
	t.Log(ind2)
	t.Log("-------------------------------------------------------------------")
	ind3, ind4 := CxESTwoPointInt(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}
