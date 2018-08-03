package inits

import (
	"math/rand"
	"testing"

	"github.com/sineatos/deag/base"
)

func getFitness() *base.Fitness {
	return base.NewFitness([]float64{-1.0})
}

func TestInitRepeat(t *testing.T) {
	genFunc := func() base.Individual {
		return base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64()}, getFitness())
	}
	individuals := InitRepeat(genFunc, 10)
	for _, ind := range individuals {
		t.Log(ind)
	}
}

func TestInitCycle(t *testing.T) {
	genIntFunc := func() base.Individual {
		return base.NewIntIndividual([]int{rand.Int(), rand.Int()}, getFitness())
	}
	genFloat64Func := func() base.Individual {
		return base.NewFloat64Individual([]float64{rand.Float64(), rand.Float64()}, getFitness())
	}
	genIntESFunc := func() base.Individual {
		return base.NewIntESIndividual([]int{rand.Int(), rand.Int()}, []float64{rand.Float64(), rand.Float64()}, getFitness())
	}
	genFloat64ESFunc := func() base.Individual {
		return base.NewFloat64ESIndividual([]float64{rand.Float64(), rand.Float64()}, []float64{rand.Float64(), rand.Float64()}, getFitness())
	}
	genBoolFunc := func() base.Individual {
		return base.NewBoolIndividual([]bool{false, true}, getFitness())
	}

	individuals := InitCycle([]func() base.Individual{
		genIntFunc,
		genFloat64Func,
		genIntESFunc,
		genFloat64ESFunc,
		genBoolFunc,
	}, 3)
	for _, ind := range individuals {
		t.Log(ind)
	}
}

func TestGenerateFloat64SliceRepeat(t *testing.T) {
	slice := GenerateFloat64SliceRepeat(rand.Float64, 10)
	t.Log(slice)
}

func TestGenerateIntSliceRepeat(t *testing.T) {
	slice := GenerateIntSliceRepeat(rand.Int, 10)
	t.Log(slice)
}

func TestGenerateBoolSliceRepeat(t *testing.T) {
	GenBool := func() bool {
		if rand.Float64() > 0.5 {
			return true
		}
		return false
	}
	slice := GenerateBoolSliceRepeat(GenBool, 10)
	t.Log(slice)
}
