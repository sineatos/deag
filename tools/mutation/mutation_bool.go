package mutation

import (
	"math/rand"

	"github.com/sineatos/deag/base"
)

// MutFlipBit flips the value of the attributes of the input individual and return the mutant.
//
// The indpb argument is the probability of each attribute to be flipped.
// This mutation is usually applied on boolean individuals.
//
// Parameters:
//
// ind: Individual to be mutated.
// indpb(float64): Independent probability for each attribute to be flipped.
//
// Returns:
// The individual which mutates with MutFlipBit
func MutFlipBit(ind *base.BoolIndividual, indpb float64) *base.BoolIndividual {
	chromosome := ind.GetChromosome().([]bool)
	for i := 0; i < ind.Len(); i++ {
		if rand.Float64() < indpb {
			chromosome[i] = !chromosome[i]
		}
	}
	return ind
}
