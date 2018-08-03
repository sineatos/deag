package base

import (
	"fmt"
)

// BoolIndividual handles bool type information
type BoolIndividual struct {
	// Chromosome
	chromosome []bool

	// Fitness
	fitness *Fitness
}

// NewBoolIndividual returns a BoolIndividual
func NewBoolIndividual(chromosome []bool, fitness *Fitness) *BoolIndividual {
	return &BoolIndividual{chromosome: chromosome, fitness: fitness}
}

// Len returns the length of the chromosome
func (ind *BoolIndividual) Len() int {
	return len(ind.chromosome)
}

// Clone returns an copy of Individual
func (ind *BoolIndividual) Clone() interface{} {
	chromosome := make([]bool, len(ind.chromosome))
	copy(chromosome, ind.chromosome)
	fitness := ind.fitness.Clone()
	return &BoolIndividual{chromosome: chromosome, fitness: fitness}
}

// GetChromosome gets the chromosome slice (not a copy)
func (ind *BoolIndividual) GetChromosome() interface{} {
	return ind.chromosome
}

// SetChromosome sets the chromosome slice
func (ind *BoolIndividual) SetChromosome(chromosome interface{}) {
	boolChromosome, ok := chromosome.([]bool)
	if !ok {
		panic(fmt.Sprintf("Chromosome should be a []int: %v", chromosome))
	}
	// if len(boolChromosome) != len(ind.chromosome) {
	// 	panic(fmt.Sprintf("Chromosome's length isn't the same with BoolChromosome: %d:%d", len(boolChromosome), len(ind.chromosome)))
	// }
	ind.chromosome = make([]bool, len(boolChromosome))
	for i, gene := range boolChromosome {
		ind.chromosome[i] = gene
	}
}

// GetFitness returns the individual's fitness (not a copy)
func (ind *BoolIndividual) GetFitness() *Fitness {
	return ind.fitness
}

// IsEqual returns if the other individual is equal to the individual
func (ind *BoolIndividual) IsEqual(other Individual) bool {
	if otherInd, ok1 := other.(*BoolIndividual); ok1 {
		if chrom, ok2 := otherInd.GetChromosome().([]bool); ok2 {
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

func (ind *BoolIndividual) String() string {
	fmtStr := "BoolIndividual{chromosome:%v, fitness:%v}"
	return fmt.Sprintf(fmtStr, ind.chromosome, ind.fitness)
}
