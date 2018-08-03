package crossover

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/tools/inits"
)

const (
	dims int = 8
)

func generateFitness() *base.Fitness {
	return base.NewFitness([]float64{-1})
}

func generateFloat64Individual() *base.Float64Individual {
	return base.NewFloat64Individual(inits.GenerateFloat64SliceRepeat(rand.Float64, dims), generateFitness())
}

func generateFloat64ESIndividual() *base.Float64ESIndividual {
	ind := inits.GenerateFloat64SliceRepeat(rand.Float64, dims)
	st := inits.GenerateFloat64SliceRepeat(rand.Float64, dims)
	return base.NewFloat64ESIndividual(ind, st, generateFitness())
}

func TestCxOnePointFloat64(t *testing.T) {
	ind1, ind2 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxOnePointFloat64(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxTwoPointFloat64(t *testing.T) {
	ind1, ind2 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxTwoPointFloat64(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxUniformFloat64(t *testing.T) {
	ind1, ind2 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxUniformFloat64(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxBlend(t *testing.T) {
	ind1, ind2 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxBlend(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxSimulatedBinary(t *testing.T) {
	ind1, ind2 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxSimulatedBinary(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxSimulatedBinaryBounded(t *testing.T) {
	low, up := 0.3, 0.6
	ind1, ind2 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxSimulatedBinaryBounded(ind1, ind2, 20, low, up)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
	t.Log("-------------------------------------------------------------------")
	lows, ups := make([]float64, dims), make([]float64, dims)
	for i := range lows {
		lows[i] = 0.3
		ups[i] = 0.8
	}
	ind5, ind6 := generateFloat64Individual(), generateFloat64Individual()
	t.Log(ind5)
	t.Log(ind6)
	ind7, ind8 := CxSimulatedBinaryBounded(ind5, ind6, 20, lows, ups)
	t.Log(ind7)
	t.Log(ind8)
	if !ind5.IsEqual(ind7) {
		t.Errorf("ind5 is not equal to ind7: %v %v", ind5, ind7)
	}
	if !ind6.IsEqual(ind8) {
		t.Errorf("ind6 is not equal to ind8: %v %v", ind6, ind8)
	}
}

func TestCxMessyOnePointFloat64(t *testing.T) {
	ind1 := base.NewFloat64Individual(inits.GenerateFloat64SliceRepeat(rand.Float64, 2), generateFitness())
	ind2 := base.NewFloat64Individual(inits.GenerateFloat64SliceRepeat(rand.Float64, 3), generateFitness())
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxMessyOnePointFloat64(ind1, ind2)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxESBlend(t *testing.T) {
	ind1, ind2 := generateFloat64ESIndividual(), generateFloat64ESIndividual()
	t.Log(ind1)
	t.Log(ind2)
	ind3, ind4 := CxESBlend(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}

func TestCxESTwoPointFloat64(t *testing.T) {
	ind1, ind2 := generateFloat64ESIndividual(), generateFloat64ESIndividual()
	t.Log(ind1)
	t.Log(ind2)
	t.Log("-------------------------------------------------------------------")
	ind3, ind4 := CxESTwoPointFloat64(ind1, ind2, 0.5)
	t.Log(ind3)
	t.Log(ind4)
	if !ind1.IsEqual(ind3) {
		t.Errorf("ind1 is not equal to ind3: %v %v", ind1, ind3)
	}
	if !ind2.IsEqual(ind4) {
		t.Errorf("ind2 is not equal to ind4: %v %v", ind2, ind4)
	}
}
