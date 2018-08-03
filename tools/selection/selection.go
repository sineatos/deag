package selection

import (
	"math"
	"math/rand"
	"sort"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/utility"
)

// SelRandom selects k individuals at random from the input individuals with replacement.
func SelRandom(individuals base.Individuals, k int) base.Individuals {
	chosen := make(base.Individuals, k)
	for i := range chosen {
		chosen[i] = individuals[rand.Intn(individuals.Len())].Clone().(base.Individual)
	}
	return chosen
}

// SelBest selects top k best individuals according to Fitness
func SelBest(individuals base.Individuals, k int) base.Individuals {
	chosen := make(base.Individuals, individuals.Len())
	copy(chosen, individuals)
	sort.Sort(sort.Reverse(chosen))
	return chosen[:k]
}

// SelWorst selects top k worst individuals according to Fitness
func SelWorst(individuals base.Individuals, k int) base.Individuals {
	chosen := make(base.Individuals, individuals.Len())
	copy(chosen, individuals)
	sort.Sort(chosen)
	return chosen[:k]
}

// SelTournament select the best individual among tournSize randomly chosen individuals, k times.
func SelTournament(individuals base.Individuals, k, tournSize int) base.Individuals {
	chosen := make(base.Individuals, k)
	for i := range chosen {
		aspirants := SelRandom(individuals, tournSize)
		maxInd := aspirants[0]
		for _, ind := range aspirants[1:] {
			if maxInd.GetFitness().Less(ind.GetFitness()) {
				maxInd = ind
			}
		}
		chosen[i] = maxInd
	}
	return chosen
}

// SelRoulette selects k individuals from the input individuals using k spins of a roulette.
// The selection is made by looking only at the first objective of each individual.
func SelRoulette(individuals base.Individuals, k int) base.Individuals {
	chosen := make(base.Individuals, k)
	sInds := SelBest(individuals, len(individuals))
	sumFits := 0.0
	for _, ind := range individuals {
		sumFits += ind.GetFitness().GetValues()[0]
	}
	for i := range chosen {
		u := rand.Float64() * sumFits
		sum := 0.0
		for _, sI := range sInds {
			sum += sI.GetFitness().GetValues()[0]
			if sum > u {
				chosen[i] = sI
				break
			}
		}
	}
	return chosen
}

// SelDoubleTournament uses the size of the individuals in order to discriminate good solutions.
// This kind of tournament is obviously useless with fixed-length representation, but has been shown to significantly reduce excessive growth of individuals, especially in GP, where it can be used as a bloat control technique (see [Luke2002fighting]_).
// This selection operator implements the double tournament technique presented in this paper.
//
// The core principle is to use a normal tournament selection, but using a special sample function to select aspirants, which is another tournament based on the size of the individuals.
// To ensure that the selection pressure is not too high, the size of the size tournament (the number of candidates evaluated) can be a real number between 1 and 2.
// In this case, the smaller individual among two will be selected with a probability size_tourn_size/2.
// For instance, if size_tourn_size is set to 1.4, then the smaller individual will have a 0.7 probability to be selected.
//
// In GP, it has been shown that this operator produces better results when it is combined with some kind of a depth limit.
//
// [Luke2002fighting] Luke and Panait, 2002, Fighting bloat with nonparametric parsimony pressure
func SelDoubleTournament(individuals base.Individuals, k int, fitnessSize, parsimonySize int, fitnessFirst bool) base.Individuals {
	if !(1 <= parsimonySize && parsimonySize <= 2) {
		panic("Parsimony tournament size has to be in the range [1, 2].")
	}

	sizeTournament := func(inds base.Individuals, kk int, selectFunc func(base.Individuals, int) base.Individuals) base.Individuals {
		chosen := make(base.Individuals, kk)
		var prob float64
		parsimonySizef64 := float64(parsimonySize)
		for i := 0; i < kk; i++ {
			prob = parsimonySizef64 / 2.0
			tmp := selectFunc(inds, 2)
			ind1, ind2 := tmp[0], tmp[1]
			if ind1.Len() > ind2.Len() {
				ind1, ind2 = ind2, ind1
			} else if ind1.Len() == ind2.Len() {
				prob = 0.5
			}

			chosen[i] = utility.If(rand.Float64() < prob, ind1, ind2).(base.Individual)
		}
		return chosen
	}

	fitTournament := func(inds base.Individuals, kk int, selectFunc func(base.Individuals, int) base.Individuals) base.Individuals {
		chosen := make(base.Individuals, kk)
		for i := 0; i < kk; i++ {
			aspirants := selectFunc(inds, fitnessSize)
			maxInd := aspirants[0]
			for _, ind := range aspirants[1:] {
				if maxInd.GetFitness().Less(ind.GetFitness()) {
					maxInd = ind
				}
			}
			chosen[i] = maxInd
		}
		return chosen
	}

	if fitnessFirst {
		tFit := func(inds base.Individuals, kk int) base.Individuals {
			return fitTournament(inds, kk, SelRandom)
		}
		return sizeTournament(individuals, k, tFit)
	}
	tSize := func(inds base.Individuals, kk int) base.Individuals {
		return sizeTournament(inds, kk, SelRandom)
	}
	return fitTournament(individuals, k, tSize)
}

// SelStochasticUniversalSampling Select the k individuals among the input individuals.
// The selection is made by using a single random value to sample all of the individuals by choosing them at evenly spaced intervals.
// The list returned contains references to the input individuals.
func SelStochasticUniversalSampling(individuals base.Individuals, k int) base.Individuals {
	chosen := make(base.Individuals, k)
	sInds := make(base.Individuals, individuals.Len())
	copy(sInds, individuals)
	sort.Sort(sort.Reverse(sInds))
	sumFits := 0.0
	for _, ind := range individuals {
		sumFits += ind.GetFitness().GetValues()[0]
	}

	distance := sumFits / float64(k)
	start := rand.Float64() * distance
	points := make([]float64, k)
	for i := range points {
		points[i] = start + float64(i)*distance
	}

	for i, p := range points {
		j := 0
		sum := sInds[j].GetFitness().GetValues()[0]
		for sum < p {
			j++
			sum += sInds[i].GetFitness().GetValues()[0]
		}
		chosen[i] = sInds[i]
	}
	return chosen
}

// SelLexicase returns an individual that does the best on the fitness cases when considered one at a time in random order.
//
// http://faculty.hampshire.edu/lspector/pubs/lexicase-IEEE-TEC.pdf
func SelLexicase(individuals base.Individuals, k int) base.Individuals {
	selectedIndividuals := make(base.Individuals, k)
	for i := 0; i < k; i++ {
		fitWeights := individuals[0].GetFitness().GetWeights()
		candidates := individuals
		cases := make([]int, len(individuals[0].GetFitness().GetValues()))
		rand.Shuffle(len(cases), func(m, n int) {
			cases[m], cases[n] = cases[n], cases[m]
		})
		for len(cases) > 0 && candidates.Len() > 1 {
			bestValForCase := candidates[0].GetFitness().GetValues()[cases[0]]
			if fitWeights[cases[0]] > 0 { // max
				for _, ind := range candidates[1:] {
					v := ind.GetFitness().GetValues()[cases[0]]
					if bestValForCase < v {
						bestValForCase = v
					}
				}
			} else { // min
				bestValForCase := candidates[0].GetFitness().GetValues()[cases[0]]
				for _, ind := range candidates[1:] {
					v := ind.GetFitness().GetValues()[cases[0]]
					if bestValForCase > v {
						bestValForCase = v
					}
				}
			}

			tmpCandidates, count := make(base.Individuals, len(candidates)), 0
			for _, ind := range candidates {
				if ind.GetFitness().GetValues()[0] == bestValForCase {
					tmpCandidates[count] = ind
					count++
				}
			}
			candidates = tmpCandidates[:count]
			cases = cases[1:]
		}
		selectedIndividuals[i] = candidates[rand.Intn(candidates.Len())]
	}
	return selectedIndividuals
}

// SelEpsilonLexicase returns an individual that does the best on the fitness cases when considered one at a time in random order.
// Requires a epsilon parameter.
//
// https://push-language.hampshire.edu/uploads/default/original/1X/35c30e47ef6323a0a949402914453f277fb1b5b0.pdf
//
// Implemented epsilon_y implementation.
func SelEpsilonLexicase(individuals base.Individuals, k int, epsilon float64) base.Individuals {
	selectedIndividuals := make(base.Individuals, k)
	for i := 0; i < k; i++ {
		fitWeights := individuals[0].GetFitness().GetWeights()
		candidates := individuals
		cases := make([]int, len(individuals[0].GetFitness().GetValues()))
		rand.Shuffle(len(cases), func(m, n int) {
			cases[m], cases[n] = cases[n], cases[m]
		})

		for len(cases) > 0 && candidates.Len() > 1 {
			bestValForCase := candidates[0].GetFitness().GetValues()[cases[0]]
			if fitWeights[cases[0]] > 0 { // max
				for _, ind := range candidates[1:] {
					v := ind.GetFitness().GetValues()[cases[0]]
					if bestValForCase < v {
						bestValForCase = v
					}
				}
				minValToSurviveCase := bestValForCase - epsilon
				tmpCandidates, count := make(base.Individuals, len(candidates)), 0
				for _, ind := range candidates {
					if ind.GetFitness().GetValues()[0] >= minValToSurviveCase {
						tmpCandidates[count] = ind
						count++
					}
				}
				candidates = tmpCandidates[:count]
			} else { // min
				bestValForCase := candidates[0].GetFitness().GetValues()[cases[0]]
				for _, ind := range candidates[1:] {
					v := ind.GetFitness().GetValues()[cases[0]]
					if bestValForCase > v {
						bestValForCase = v
					}
				}
				maxValToSurviveCase := bestValForCase + epsilon
				tmpCandidates, count := make(base.Individuals, len(candidates)), 0
				for _, ind := range candidates {
					if ind.GetFitness().GetValues()[0] <= maxValToSurviveCase {
						tmpCandidates[count] = ind
						count++
					}
				}
				candidates = tmpCandidates[:count]
			}
			cases = cases[1:]
		}
		selectedIndividuals[i] = candidates[rand.Intn(candidates.Len())]
	}
	return selectedIndividuals
}

// SelAutomaticEpsilonLexicase returns an individual that does the best on the fitness cases when considered one at a time in random order.
//
// https://push-language.hampshire.edu/uploads/default/original/1X/35c30e47ef6323a0a949402914453f277fb1b5b0.pdf
//
// Implemented lambda_epsilon_y implementation.
func SelAutomaticEpsilonLexicase(individuals base.Individuals, k int) base.Individuals {
	selectedIndividuals := make(base.Individuals, k)
	for i := 0; i < k; i++ {
		fitWeights := individuals[0].GetFitness().GetWeights()
		candidates := individuals
		cases := make([]int, len(individuals[0].GetFitness().GetValues()))
		rand.Shuffle(len(cases), func(m, n int) {
			cases[m], cases[n] = cases[n], cases[m]
		})

		for len(cases) > 0 && candidates.Len() > 1 {

			errorsForThisCase := make([]float64, candidates.Len())
			errorsForThisCaseAbsolute := make([]float64, candidates.Len())
			for i, ind := range candidates {
				errorsForThisCase[i] = ind.GetFitness().GetValues()[cases[0]]
			}
			medianVal := median(errorsForThisCase)

			for i, val := range errorsForThisCase {
				errorsForThisCaseAbsolute[i] = math.Abs(medianVal - val)
			}
			medianAbsoluteDeviation := median(errorsForThisCaseAbsolute)

			bestValForCase := candidates[0].GetFitness().GetValues()[cases[0]]
			if fitWeights[cases[0]] > 0 { // max
				for _, ind := range candidates[1:] {
					v := ind.GetFitness().GetValues()[cases[0]]
					if bestValForCase < v {
						bestValForCase = v
					}
				}
				minValToSurviveCase := bestValForCase - medianAbsoluteDeviation
				tmpCandidates, count := make(base.Individuals, len(candidates)), 0
				for _, ind := range candidates {
					if ind.GetFitness().GetValues()[0] >= minValToSurviveCase {
						tmpCandidates[count] = ind
						count++
					}
				}
				candidates = tmpCandidates[:count]
			} else { // min
				for _, ind := range candidates[1:] {
					v := ind.GetFitness().GetValues()[cases[0]]
					if bestValForCase > v {
						bestValForCase = v
					}
				}
				maxValToSurviveCase := bestValForCase + medianAbsoluteDeviation
				tmpCandidates, count := make(base.Individuals, len(candidates)), 0
				for _, ind := range candidates {
					if ind.GetFitness().GetValues()[0] <= maxValToSurviveCase {
						tmpCandidates[count] = ind
						count++
					}
				}
				candidates = tmpCandidates[:count]
			}
			cases = cases[1:]
		}
		selectedIndividuals[i] = candidates[rand.Intn(candidates.Len())]
	}
	return selectedIndividuals
}

// median sorts the slice and returns the median
func median(arr []float64) float64 {
	sort.Float64s(arr)
	l := len(arr)
	if l&1 == 1 {
		return arr[(l-1)/2]
	}
	return (arr[l/2] + arr[l/2+1]) / 2.0
}
