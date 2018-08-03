package mutation

import (
	"math/rand"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/utility"
)

// MutUniformInt mutates an individual by replacing attributes, with probability indpb, by a integer uniformly drawn between low and up inclusively.
//
// Paramters:
//
// ind: Individual to be mutated.
//
// low(int or []int): The lower bound of the range from wich to draw the new integer.
//
// up(int or []int): The upper bound of the range from wich to draw the new integer.
//
// indpb(float64): Independent probability for each attribute to be mutated.
//
// Returns:
//
// The individual which mutates with MutUniformInt
func MutUniformInt(ind *base.IntIndividual, low interface{}, up interface{}, indpb float64) *base.IntIndividual {
	size := ind.Len()
	lows := utility.Interface2IntSlice("low", low, size)
	ups := utility.Interface2IntSlice("up", up, size)
	chromosome := ind.GetChromosome().([]int)
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			chromosome[i] = lows[i] + rand.Intn(ups[i]-lows[i])
		}
	}
	return ind
}
