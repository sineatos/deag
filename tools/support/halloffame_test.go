package support

import (
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
)

func TestDefaultHallOfFame(t *testing.T) {
	hof1 := NewDefaultHallOfFame(2, nil)
	ind1 := base.NewFloat64Individual([]float64{0.0, 0.0, 0.0}, base.NewFitness([]float64{-1.0}))
	cInd1 := ind1.Clone().(*base.Float64Individual)
	ans1, cAns1 := benchmarks.Ackley(ind1), benchmarks.Ackley(cInd1)
	ind1.GetFitness().SetValues(ans1)
	cInd1.GetFitness().SetValues(cAns1)
	if !ind1.IsEqual(cInd1) {
		t.Errorf("individuals aren't the same :%v %v", ind1, cInd1)
	}
	ind2 := base.NewFloat64Individual([]float64{2.0, 2.0, 2.0}, base.NewFitness([]float64{-1.0}))
	ans2 := benchmarks.Ackley(ind2)
	ind2.GetFitness().SetValues(ans2)
	ind3 := base.NewFloat64Individual([]float64{1.0, 1.0, 1.0}, base.NewFitness([]float64{-1.0}))
	ans3 := benchmarks.Ackley(ind3)
	ind3.GetFitness().SetValues(ans3)
	if hof1.Len() != 0 {
		t.Errorf("hof1's len != 0: %v", hof1.Len())
	}
	hof1.Update(base.Individuals{ind1})
	if hof1.Len() != 1 {
		t.Errorf("hof1's len != 1: %v", hof1.Len())
	}
	t.Log(hof1)
	hof1.Update(base.Individuals{cInd1})
	if hof1.Len() != 1 {
		t.Errorf("hof1's len != 1: %v", hof1.Len())
	}
	t.Log(hof1)
	hof1.Update(base.Individuals{ind2})
	if hof1.Len() != 2 {
		t.Errorf("hof1's len != 2: %v", hof1.Len())
	}
	t.Log(hof1)
	t.Log("Insert ind3 ---------------------------------------------------")
	hof1.Insert(ind3)
	t.Log(hof1)
	t.Log("Remove(0) ---------------------------------------------------")
	hof1.Remove(0)
	t.Log(hof1)
	t.Log("Insert ind2 ---------------------------------------------------")
	hof1.Insert(ind2)
	t.Log(hof1)
	t.Log("Reversed() ---------------------------------------------------")
	t.Log(hof1.Reversed())
	t.Log("Get(0) ---------------------------------------------------")
	t.Log(hof1.Get(0))
	t.Log("Clear() ---------------------------------------------------")
	hof1.Clear()
	t.Log(hof1)
}
