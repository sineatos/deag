package benchmarks

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/tools/inits"
	"github.com/sineatos/deag/utility"
)

const (
	bdims int = 40
)

func generateFitness() *base.Fitness {
	return base.NewFitness([]float64{-1})
}

func genBoolIndividual() *base.BoolIndividual {
	f := func() bool { return utility.If(rand.Float64() > 0.5, true, false).(bool) }
	return base.NewBoolIndividual(inits.GenerateBoolSliceRepeat(f, bdims), generateFitness())
}

func TestChuangF1(t *testing.T) {
	b1, b2 := make([]bool, bdims), make([]bool, bdims)
	for i := range b1 {
		b1[i] = true
		b2[i] = false
	}
	bInd1 := base.NewBoolIndividual(b1, generateFitness())
	bInd2 := base.NewBoolIndividual(b2, generateFitness())
	ind1 := genBoolIndividual()
	ind2 := genBoolIndividual()
	ans1, ans2 := ChuangF1(ind1), ChuangF1(ind2)
	ind1.GetFitness().SetValues(ans1)
	ind2.GetFitness().SetValues(ans2)
	t.Log(ind1)
	t.Log(ind2)
	t.Log("-------------------------------------------------------------------------------")
	bAns1, bAns2 := ChuangF1(bInd1), ChuangF1(bInd2)
	bInd1.GetFitness().SetValues(bAns1)
	bInd2.GetFitness().SetValues(bAns2)
	t.Log(bInd1)
	t.Log(bInd2)
}

func TestChuangF2(t *testing.T) {
	b1, b2 := make([]bool, bdims), make([]bool, bdims)
	for i := range b1 {
		b1[i] = true
		b2[i] = false
	}
	bInd1 := base.NewBoolIndividual(b1, generateFitness())
	bInd2 := base.NewBoolIndividual(b2, generateFitness())
	ind1 := genBoolIndividual()
	ind2 := genBoolIndividual()
	ans1, ans2 := ChuangF2(ind1), ChuangF2(ind2)
	ind1.GetFitness().SetValues(ans1)
	ind2.GetFitness().SetValues(ans2)
	t.Log(ind1)
	t.Log(ind2)
	t.Log("-------------------------------------------------------------------------------")
	bAns1, bAns2 := ChuangF2(bInd1), ChuangF2(bInd2)
	bInd1.GetFitness().SetValues(bAns1)
	bInd2.GetFitness().SetValues(bAns2)
	t.Log(bInd1)
	t.Log(bInd2)
}

func TestChuangF3(t *testing.T) {
	b1, b2 := make([]bool, bdims), make([]bool, bdims)
	for i := range b1 {
		b1[i] = true
		b2[i] = false
	}
	bInd1 := base.NewBoolIndividual(b1, generateFitness())
	bInd2 := base.NewBoolIndividual(b2, generateFitness())
	ind1 := genBoolIndividual()
	ind2 := genBoolIndividual()
	ans1, ans2 := ChuangF3(ind1), ChuangF3(ind2)
	ind1.GetFitness().SetValues(ans1)
	ind2.GetFitness().SetValues(ans2)
	t.Log(ind1)
	t.Log(ind2)
	t.Log("-------------------------------------------------------------------------------")
	bAns1, bAns2 := ChuangF3(bInd1), ChuangF3(bInd2)
	bInd1.GetFitness().SetValues(bAns1)
	bInd2.GetFitness().SetValues(bAns2)
	t.Log(bInd1)
	t.Log(bInd2)
}

func TestRoyalRoad1(t *testing.T) {
	b1, b2 := make([]bool, bdims), make([]bool, bdims)
	for i := range b1 {
		b1[i] = true
		b2[i] = false
	}
	ind1 := genBoolIndividual()
	ind2 := genBoolIndividual()
	ans1, ans2 := RoyalRoad1(ind1, 1), RoyalRoad1(ind2, 1)
	ind1.GetFitness().SetValues(ans1)
	ind2.GetFitness().SetValues(ans2)
	t.Log(ind1)
	t.Log(ind2)
}

func TestRoyalRoad2(t *testing.T) {
	b1, b2 := make([]bool, bdims), make([]bool, bdims)
	for i := range b1 {
		b1[i] = true
		b2[i] = false
	}
	ind1 := genBoolIndividual()
	ind2 := genBoolIndividual()
	ans1, ans2 := RoyalRoad2(ind1, 1), RoyalRoad2(ind2, 1)
	ind1.GetFitness().SetValues(ans1)
	ind2.GetFitness().SetValues(ans2)
	t.Log(ind1)
	t.Log(ind2)
}
