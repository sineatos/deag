package selection

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
	"github.com/sineatos/deag/tools/inits"
)

const (
	dims    int = 10
	amounts int = 100
)

func generateFitness() *base.Fitness {
	return base.NewFitness([]float64{-1})
}

func generateFloat64Individual() base.Individual {
	f := func() float64 { return rand.Float64() * 100.0 }
	return base.NewFloat64Individual(inits.GenerateFloat64SliceRepeat(f, dims), generateFitness())
}

func genIndividuals() base.Individuals {
	individuals := inits.InitRepeat(generateFloat64Individual, amounts)
	for _, ind := range individuals {
		cInd, _ := ind.(*base.Float64Individual)
		val := benchmarks.Ackley(cInd)
		ind.GetFitness().SetValues(val)
	}
	return individuals
}

func TestSelRandom(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelRandom(individuals, 2)
	t.Log(chosen)
}

func TestSelBest(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelBest(individuals, 4)
	t.Log(chosen)
}

func TestSelWorst(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelWorst(individuals, 2)
	t.Log(chosen)
}

func TestSelTournament(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelTournament(individuals, 3, 2)
	t.Log(chosen)
}

func TestRoulette(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelRoulette(individuals, 3)
	t.Log(chosen)
}

func TestSelDoubleTournament(t *testing.T) {
	individuals1 := genIndividuals()
	t.Log(individuals1)
	chosen1 := SelDoubleTournament(individuals1, 3, 1, 1, false)
	t.Log(chosen1)
	t.Log("------------------------------------------------------------")
	individuals2 := genIndividuals()
	t.Log(individuals2)
	chosen2 := SelDoubleTournament(individuals2, 3, 1, 1, true)
	t.Log(chosen2)
}

func TestSelStochasticUniversalSampling(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelStochasticUniversalSampling(individuals, 3)
	t.Log(chosen)
}

func TestSelLexicase(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelLexicase(individuals, 3)
	t.Log(chosen)
}

func TestSelEpsilonLexicase(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelEpsilonLexicase(individuals, 3, 0.6)
	t.Log(chosen)
}

func TestSelAutomaticEpsilonLexicase(t *testing.T) {
	individuals := genIndividuals()
	t.Log(individuals)
	chosen := SelAutomaticEpsilonLexicase(individuals, 3)
	t.Log(chosen)
}
