package emo

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
	"github.com/sineatos/deag/tools/crossover"
	"github.com/sineatos/deag/tools/inits"
	"github.com/sineatos/deag/tools/mutation"
	"github.com/sineatos/deag/tools/support"
)

func newNSGA2Statistics() support.Statistics {
	mStat := support.NewDefaultMultiStatistics("MultiStatistics")
	indStat := support.NewStatisticsBasedOnMOFitness("")
	indStat.Register("min", support.StatMOFitnessMin)
	indStat.Register("max", support.StatMOFitnessMax)
	mStat.AddStats(indStat)
	return mStat
}

func newNSGA2Population(size, dims int, low, up float64) base.Individuals {
	pop := make(base.Individuals, size)
	fit := base.NewFitness([]float64{-1.0, -1.0})
	limit := func() float64 {
		return low + rand.Float64()*(up-low)
	}
	getData := func() []float64 { return inits.GenerateFloat64SliceRepeat(limit, dims) }
	for i := range pop {
		pop[i] = base.NewFloat64Individual(getData(), fit.Clone())
	}
	return pop
}

func TestNSGA2(t *testing.T) {
	ngen, dims, size, cxpb, eta := 250, 30, 100, 0.9, 20.0
	boundLow, boundUp := 0.0, 1.0
	stat := newNSGA2Statistics()
	pop := newNSGA2Population(size, dims, boundLow, boundUp)
	logbook := support.NewDefaultLogbook(0, 0)

	log := func(g int) {
		record := stat.Compile(pop)
		record[support.GEN] = g
		record[support.FES] = len(pop)
		logbook.Record(record)
	}

	// mate, mutate, select and benchmark
	mate := func(ind1, ind2 base.Individual) (*base.Float64Individual, *base.Float64Individual) {
		fInd1, fInd2 := ind1.(*base.Float64Individual), ind2.(*base.Float64Individual)
		return crossover.CxSimulatedBinaryBounded(fInd1, fInd2, eta, boundLow, boundUp)
	}
	mutate := func(ind base.Individual) *base.Float64Individual {
		fInd := ind.(*base.Float64Individual)
		return mutation.MutPolyNomialBounded(fInd, eta, boundLow, boundUp, 1.0/float64(dims))
	}
	zelect := SelNSGA2
	benchmark := benchmarks.ZDT1

	for _, ind := range pop {
		fInd := ind.(*base.Float64Individual)
		ind.GetFitness().SetValues(benchmark(fInd))
	}
	log(0)

	for gen := 1; gen < ngen; gen++ {
		offspring := make(base.Individuals, 0, size)
		for _, ind := range pop {
			offspring = append(offspring, ind.Clone().(base.Individual))
		}

		for i, j := 0, 1; i < size; {
			if rand.Float64() <= cxpb {
				mate(offspring[i], offspring[j])
			}
			ind1 := mutate(offspring[i])
			ind2 := mutate(offspring[j])
			ind1.GetFitness().SetValues(benchmark(ind1))
			ind2.GetFitness().SetValues(benchmark(ind2))
			i += 2
			j += 2
		}

		tempPop := make(base.Individuals, 0, size*2)
		tempPop = append(tempPop, pop...)
		tempPop = append(tempPop, offspring...)
		pop = zelect(tempPop, size)
		log(gen)
	}

	t.Log(logbook.String())
}
