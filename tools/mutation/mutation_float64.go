package mutation

import (
	"math"
	"math/rand"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/utility"
)

/***************************
 * GA Mutations            *
 ***************************/

// MutGaussian is a function applies a gaussian mutation of mean and standard deviation on the input individual.
//
// This mutation expects a sequence individual composed of real valued attributes.
// The indpb argument is the probability of each attribute to be mutated.
//
// Parameters:
//
// ind: Individual to be mutated.
//
// mu(float64 or []float64): Mean or sequence of means for the gaussian addition mutation.
//
// sigma(float64 or []float64): Standard deviation or sequence of standard deviations for the gaussian addition mutation.
//
// indpb(float64): Independent probability for each attribute to be mutated.
//
// Returns:
//
// The individual which mutates with MutGaussian
func MutGaussian(ind *base.Float64Individual, mu interface{}, sigma interface{}, indpb float64) *base.Float64Individual {
	size := ind.Len()
	mus := utility.Interface2Float64Slice("mu", mu, size)
	sigmas := utility.Interface2Float64Slice("sigma", sigma, size)
	chromosome := ind.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			chromosome[i] += rand.NormFloat64()*sigmas[i] + mus[i]
		}
	}
	return ind
}

// MutPolyNomialBounded is the polynomial mutation as implemented in original NSGA-II algorithm in C by Deb.
//
// Parameters:
//
// ind: Individual to be mutated.
//
// eta(float64): Crowding degree of the mutation. A high eta will produce a mutant resembling its parent, while a small eta will produce a solution much more different.
//
// low(float64 or []float64): The lower bound of the search space.
//
// up(float64 or []float64): The upper bound of the search space.
//
// Returns:
//
// The individual which mutates with MutPolyNomialBounded
func MutPolyNomialBounded(ind *base.Float64Individual, eta float64, low interface{}, up interface{}, indpb float64) *base.Float64Individual {
	size := ind.Len()
	lows := utility.Interface2Float64Slice("low", low, size)
	ups := utility.Interface2Float64Slice("up", up, size)
	chromosome := ind.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		if rand.Float64() <= indpb {
			x := chromosome[i]
			delta1 := (x - lows[i]) / (ups[i] - lows[i])
			delta2 := (ups[i] - x) / (ups[i] - lows[i])
			ran := rand.Float64()
			mutPow := 1.0 / (eta + 1.0)
			var deltaQ float64

			if ran < 0.5 {
				xy := 1.0 - delta1
				val := 2.0*ran + (1.0-2.0*ran)*math.Pow(xy, eta+1.0)
				deltaQ = math.Pow(val, mutPow) - 1.0
			} else {
				xy := 1.0 - delta2
				val := 2.0*(1.0-ran) + 2.0*(ran-0.5)*math.Pow(xy, eta+1.0)
				deltaQ = 1.0 - math.Pow(val, mutPow)
			}

			x = x + deltaQ*(ups[i]-lows[i])
			x = math.Min(math.Max(x, lows[i]), ups[i])
			chromosome[i] = x
		}
	}
	return ind
}

// MutShuffleIndexesFloat64 shuffles the attributes of the input individual and return the mutant.
//
// The indpb argument is the probability of each attribute to be moved.
// Usually this mutation is applied on vector of indices.
//
// Parameters:
//
// ind: Individual to be mutated.
// indpb(float64): Independent probability for each attribute to be exchanged to another position.
//
// Returns:
//
// The individual which mutates with MutShuffleIndexes
func MutShuffleIndexesFloat64(ind *base.Float64Individual, indpb float64) *base.Float64Individual {
	size := ind.Len()
	chromosome := ind.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			swapIndex := rand.Intn(size - 2)
			if swapIndex >= i {
				swapIndex += i
			}
			chromosome[swapIndex], chromosome[i] = chromosome[i], chromosome[swapIndex]
		}
	}
	return ind
}

/***************************
 * ES Mutations            *
 ***************************/

// MutESLogNormal mutates an evolution strategy according to its strategy attribute as described in [Beyer2002].
//
// First the strategy is mutated according to an extended log normal rule,
// 	:math:`\\boldsymbol{\sigma}_t = \\exp(\\tau_0 \mathcal{N}_0(0, 1)) \\left[ \\sigma_{t-1, 1}\\exp(\\tau \mathcal{N}_1(0, 1)), \ldots, \\sigma_{t-1, n} \\exp(\\tau \mathcal{N}_n(0, 1))\\right]`, with :math:`\\tau_0 = \\frac{c}{\\sqrt{2n}}` and :math:`\\tau = \\frac{c}{\\sqrt{2\\sqrt{n}}}`, the the individual is mutated by a normal distribution of mean 0 and standard deviation of :math:`\\boldsymbol{\sigma}_{t}`
// (its current strategy) then .
// A recommended choice is c=1 when using a evolution strategy [Beyer2002] [Schwefel1995]
//
// [Beyer2002] Beyer and Schwefel, 2002, Evolution strategies - A Comprehensive Introduction
//
// [Schwefel1995] Schwefel, 1995, Evolution and Optimum Seeking. Wiley, New York, NY
//
// Parameters:
// ind: individual to be mutated.
//
// c(float64): The learning parameter.
//
// indpb(float64): Independent probability for each attribute to be mutated.
//
// Returns:
//
// The individual which mutates with MutESLogNormal
func MutESLogNormal(ind *base.Float64ESIndividual, c float64, indpb float64) *base.Float64ESIndividual {
	size := ind.Len()
	t := c / math.Sqrt(2.0*math.Sqrt(float64(size)))
	t0 := c / math.Sqrt(2.0*float64(size))
	n := rand.NormFloat64()
	t0n := t0 * n
	chromosome := ind.GetChromosome().([]float64)
	strategies := ind.GetStrategies()
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			strategies[i] = strategies[i] * math.Exp(t0n+t*rand.NormFloat64())
			chromosome[i] = chromosome[i] + strategies[i]*rand.NormFloat64()
		}
	}
	return ind
}
