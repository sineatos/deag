package base

import (
	"fmt"
	"math"
)

// Float64Individual handles float64 type information
type Float64Individual struct {
	// Chromosome
	chromosome []float64

	// Fitness
	fitness *Fitness
}

// NewFloat64Individual returns a Float64Individual
func NewFloat64Individual(chromosome []float64, fitness *Fitness) *Float64Individual {
	return &Float64Individual{chromosome: chromosome, fitness: fitness}
}

// Len returns the length of the chromosome
func (ind *Float64Individual) Len() int {
	return len(ind.chromosome)
}

// Clone returns an copy of Individual
func (ind *Float64Individual) Clone() interface{} {
	chromosome := make([]float64, len(ind.chromosome))
	copy(chromosome, ind.chromosome)
	fitness := ind.fitness.Clone()
	return &Float64Individual{chromosome: chromosome, fitness: fitness}
}

// GetChromosome gets the chromosome slice (not a copy)
func (ind *Float64Individual) GetChromosome() interface{} {
	return ind.chromosome
}

// SetChromosome sets the chromosome slice
func (ind *Float64Individual) SetChromosome(chromosome interface{}) {
	float64Chromosome, ok := chromosome.([]float64)
	if !ok {
		panic(fmt.Sprintf("Chromosome should be a []float64: %v", chromosome))
	}
	// if len(float64Chromosome) != len(ind.chromosome) {
	// 	panic(fmt.Sprintf("Chromosome's length isn't the same with Float64Individual: %d:%d", len(float64Chromosome), len(ind.chromosome)))
	// }
	ind.chromosome = make([]float64, len(float64Chromosome))
	for i, gene := range float64Chromosome {
		ind.chromosome[i] = gene
	}
}

// GetFitness returns the individual's fitness (not a copy)
func (ind *Float64Individual) GetFitness() *Fitness {
	return ind.fitness
}

// IsEqual returns if the other individual is equal to the individual
func (ind *Float64Individual) IsEqual(other Individual) bool {
	if otherInd, ok1 := other.(*Float64Individual); ok1 {
		if chrom, ok2 := otherInd.GetChromosome().([]float64); ok2 {
			if len(chrom) == len(ind.chromosome) {
				for i, c := range ind.chromosome {
					if math.Abs(chrom[i]-c) > 1E-14 {
						return false
					}
				}
				return true
			}
		}
	}
	return false
}

func (ind *Float64Individual) String() string {
	fmtStr := "Float64Individual{chromosome:%v, fitness:%v}"
	return fmt.Sprintf(fmtStr, ind.chromosome, ind.fitness)
}

// Float64ESIndividual handles float64 type information
type Float64ESIndividual struct {
	// Float64Individual
	Float64Individual

	// strategies
	strategies []float64
}

// NewFloat64ESIndividual returns a Float64ESIndividual
func NewFloat64ESIndividual(chromosome, strategies []float64, fitness *Fitness) *Float64ESIndividual {
	return &Float64ESIndividual{Float64Individual: *NewFloat64Individual(chromosome, fitness), strategies: strategies}
}

// SLen returns the size of strategies
func (ind *Float64ESIndividual) SLen() int {
	return len(ind.strategies)
}

// Clone returns an copy of Individual
func (ind *Float64ESIndividual) Clone() interface{} {
	f64Ind := ind.Float64Individual.Clone().(*Float64Individual)
	strategies := make([]float64, len(ind.strategies))
	copy(strategies, ind.strategies)
	return &Float64ESIndividual{Float64Individual: *f64Ind, strategies: strategies}
}

// GetStrategies return strategies(not copy)
func (ind *Float64ESIndividual) GetStrategies() []float64 {
	return ind.strategies
}

// SetStrategies set strategies
func (ind *Float64ESIndividual) SetStrategies(strategies []float64) {
	ind.strategies = make([]float64, len(strategies))
	for i, strategy := range strategies {
		ind.strategies[i] = strategy
	}
}

// IsEqual returns if the other individual is equal to the individual
func (ind *Float64ESIndividual) IsEqual(other Individual) bool {
	if otherInd, ok1 := other.(*Float64ESIndividual); ok1 {
		if chrom, ok2 := otherInd.GetChromosome().([]float64); ok2 {
			if len(chrom) == len(ind.chromosome) {
				for i, c := range ind.chromosome {
					if math.Abs(chrom[i]-c) > 1E-14 {
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

func (ind *Float64ESIndividual) String() string {
	fmtStr := "Float64ESIndividual{chromosome:%v, fitness:%v, strategies:%v}"
	return fmt.Sprintf(fmtStr, ind.chromosome, ind.fitness, ind.strategies)
}
