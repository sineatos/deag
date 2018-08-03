package pso

import (
	"math"
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
	"github.com/sineatos/deag/tools/constraint"
	"github.com/sineatos/deag/tools/inits"
	"github.com/sineatos/deag/tools/support"
)

func newPSOStatistics() support.Statistics {
	mStat := support.NewDefaultMultiStatistics("MultiStatistics")
	indStat := support.NewStatisticsBasedOnFitness("fitness")
	indStat.Register("min", support.StatFitnessMin)
	indStat.Register("max", support.StatFitnessMax)
	indStat.Register("avg", support.StatFitnessAvg)
	indStat.Register("std", support.StatFitnessStd)
	mStat.AddStats(indStat)
	return mStat
}

func newPSOConstraint(low, up float64, evaluator benchmarks.Float64Evaluator) func(*base.Float64Individual) []float64 {
	isFeasible := func(individual *base.Float64Individual) bool {
		chrom := individual.GetChromosome().([]float64)
		for _, c := range chrom {
			if c < low || c > up {
				return false
			}
		}
		return true
	}
	adjust := func(individual *base.Float64Individual) *base.Float64Individual {
		chrom := individual.GetChromosome().([]float64)
		for i, c := range chrom {
			if c < low || c > up {
				chrom[i] = low + (up-low)*rand.Float64()
			}
		}
		return individual
	}
	con := constraint.NewClosestValidPenalty(isFeasible, adjust, 0.0, nil, evaluator)
	wrapper := func(individual *base.Float64Individual) []float64 {
		return con.AdjustAndEvolve(individual)
	}
	return wrapper
}

func newPSOPopulation(size, dims int, low, up, smin, smax float64) base.Individuals {
	pop := make(base.Individuals, size)
	fit := base.NewFitness([]float64{-1.0})
	limit := func() float64 {
		return low + rand.Float64()*(up-low)
	}
	sLimit := func() float64 {
		return smin + rand.Float64()*(smax-smin)
	}
	getData := func() []float64 { return inits.GenerateFloat64SliceRepeat(limit, dims) }
	getSpeed := func() []float64 { return inits.GenerateFloat64SliceRepeat(sLimit, dims) }
	for i := range pop {
		pop[i] = NewParticle(getData(), getSpeed(), smin, smax, fit.Clone())
	}
	return pop
}

func TestPSO(t *testing.T) {
	c1, c2 := 2.0, 2.0
	maxGen, maxFES := 200, math.MaxInt64
	st, hof := newPSOStatistics(), support.NewDefaultHallOfFame(2, nil)
	size, dims := 5, 2
	low, up := -6.0, 6.0
	smin, smax := -3.0, 3.0
	pop := newPSOPopulation(size, dims, low, up, smin, smax)
	evaluator := newPSOConstraint(low, up, benchmarks.H1)
	evol := NewPSO(c1, c2, maxGen, maxFES, st, hof, evaluator)
	evol.Init(pop)
	evol.Run()
	logbook := evol.GetLogbook()
	hof1 := evol.GetHallOfFame()
	if hof != hof1 {
		t.Error("hof != hof1")
	}
	t.Log("\n" + logbook.String())
	for i := 0; i < hof.Len(); i++ {
		t.Log(hof.Get(i))
	}
}
