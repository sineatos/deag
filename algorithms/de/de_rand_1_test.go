package de

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

func newDERand1Statistics() support.Statistics {
	mStat := support.NewDefaultMultiStatistics("MultiStatistics")
	indStat := support.NewStatisticsBasedOnFitness("fitness")
	indStat.Register("min", support.StatFitnessMin)
	indStat.Register("max", support.StatFitnessMax)
	indStat.Register("avg", support.StatFitnessAvg)
	indStat.Register("std", support.StatFitnessStd)
	mStat.AddStats(indStat)
	return mStat
}

func newDERand1Constraint(low, up float64, evaluator benchmarks.Float64Evaluator) func(*base.Float64Individual) []float64 {
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

func newDERand1Population(size, dims int, low, up float64) base.Individuals {
	pop := make(base.Individuals, size)
	fit := base.NewFitness([]float64{-1.0})
	limit := func() float64 {
		return low + rand.Float64()*(up-low)
	}
	getData := func() []float64 { return inits.GenerateFloat64SliceRepeat(limit, dims) }
	for i := range pop {
		pop[i] = base.NewFloat64Individual(getData(), fit.Clone())
	}
	return pop
}

func TestDERand1(t *testing.T) {
	f, cr := 0.5, 0.5
	maxGen, maxFES := 200, math.MaxInt64
	st, hof := newDEBest1Statistics(), support.NewDefaultHallOfFame(3, nil)
	size, dims := 200, 5
	low, up := -5.12, 5.12
	pop := newDEBest1Population(size, dims, low, up)
	evaluator := newDEBest1Constraint(low, up, benchmarks.Rastrigin)
	evol := NewDERand1(f, cr, maxGen, maxFES, st, hof, evaluator)
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
