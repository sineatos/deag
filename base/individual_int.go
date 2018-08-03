package base

import (
	"fmt"
)

// IntIndividual handles integer type information
type IntIndividual struct {
	// Chromosome
	chromosome []int

	// Fitness
	fitness *Fitness
}

// NewIntIndividual returns a IntIndividual
func NewIntIndividual(chromosome []int, fitness *Fitness) *IntIndividual {
	return &IntIndividual{chromosome: chromosome, fitness: fitness}
}

// Len returns the length of the chromosome
func (ind *IntIndividual) Len() int {
	return len(ind.chromosome)
}

// Clone returns an copy of Individual
func (ind *IntIndividual) Clone() interface{} {
	chromosome := make([]int, len(ind.chromosome))
	copy(chromosome, ind.chromosome)
	fitness := ind.fitness.Clone()
	return &IntIndividual{chromosome: chromosome, fitness: fitness}
}

// GetChromosome gets the chromosome slice (not a copy)
func (ind *IntIndividual) GetChromosome() interface{} {
	return ind.chromosome
}

// SetChromosome sets the chromosome slice
func (ind *IntIndividual) SetChromosome(chromosome interface{}) {
	intChromosome, ok := chromosome.([]int)
	if !ok {
		panic(fmt.Sprintf("Chromosome should be a []int: %v", chromosome))
	}
	// if len(intChromosome) != len(ind.chromosome) {
	// 	panic(fmt.Sprintf("Chromosome's length isn't the same with IntIndividual: %d:%d", len(intChromosome), len(ind.chromosome)))
	// }
	ind.chromosome = make([]int, len(intChromosome))
	for i, gene := range intChromosome {
		ind.chromosome[i] = gene
	}
}

// GetFitness returns the individual's fitness (not a copy)
func (ind *IntIndividual) GetFitness() *Fitness {
	return ind.fitness
}

// IsEqual returns if the other individual is equal to the individual
func (ind *IntIndividual) IsEqual(other Individual) bool {
	if otherInd, ok1 := other.(*IntIndividual); ok1 {
		if chrom, ok2 := otherInd.GetChromosome().([]int); ok2 {
			if len(chrom) == len(ind.chromosome) {
				for i, c := range ind.chromosome {
					if chrom[i] != c {
						return false
					}
				}
				return true
			}
		}
	}
	return false
}

func (ind *IntIndividual) String() string {
	fmtStr := "IntIndividual{chromosome:%v, fitness:%v}"
	return fmt.Sprintf(fmtStr, ind.chromosome, ind.fitness)
}

// IntESIndividual handles int type information
type IntESIndividual struct {
	// IntESIndividual
	IntIndividual

	// strategies
	strategies []float64
}

// NewIntESIndividual returns a IntESIndividual
func NewIntESIndividual(chromosome []int, strategies []float64, fitness *Fitness) *IntESIndividual {
	return &IntESIndividual{IntIndividual: *NewIntIndividual(chromosome, fitness), strategies: strategies}
}

// SLen returns the size of strategies
func (ind *IntESIndividual) SLen() int {
	return len(ind.strategies)
}

// Clone returns an copy of Individual
func (ind *IntESIndividual) Clone() interface{} {
	intInd := ind.IntIndividual.Clone().(*IntIndividual)
	strategies := make([]float64, len(ind.strategies))
	copy(strategies, ind.strategies)
	return &IntESIndividual{IntIndividual: *intInd, strategies: strategies}
}

// GetStrategies return strategies(not copy)
func (ind *IntESIndividual) GetStrategies() []float64 {
	return ind.strategies
}

// SetStrategies set strategies
func (ind *IntESIndividual) SetStrategies(strategies []float64) {
	ind.strategies = make([]float64, len(strategies))
	for i, strategy := range strategies {
		ind.strategies[i] = strategy
	}
}

// IsEqual returns if the other individual is equal to the individual
func (ind *IntESIndividual) IsEqual(other Individual) bool {
	if otherInd, ok1 := other.(*IntESIndividual); ok1 {
		if chrom, ok2 := otherInd.GetChromosome().([]int); ok2 {
			if len(chrom) == len(ind.chromosome) {
				for i, c := range ind.chromosome {
					if chrom[i] != c {
						return false
					}
				}
				// otherStrategies := otherInd.GetStrategies()
				// if len(otherStrategies) == len(ind.strategies) {
				// 	for i, s := range ind.strategies {
				// 		if otherStrategies[i] != s {
				// 			return false
				// 		}
				// 	}
				// 	return true
				// }
			}
		}
	}
	return true
	// return false
}

func (ind *IntESIndividual) String() string {
	fmtStr := "IntESIndividual{chromosome:%v, fitness:%v, strategies:%v}"
	return fmt.Sprintf(fmtStr, ind.chromosome, ind.fitness, ind.strategies)
}
