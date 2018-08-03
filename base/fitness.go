package base

import (
	"fmt"
)

// Fitness is a measure of quality of a solution.
// Define Fitness A greater than B means A better than B, A less than B means A worse than B
type Fitness struct {
	weights []float64
	wvalues []float64
	values  []float64
	valid   bool
}

// NewFitness returns a fitness value using weights and the values is a zero vector
func NewFitness(weights []float64) *Fitness {
	if len(weights) > 0 {
		fitness := &Fitness{weights: weights, valid: false}
		return fitness
	}
	return nil
}

// NewFitnessWithValues is used to get a fitness value
func NewFitnessWithValues(weights []float64, values []float64) *Fitness {
	if len(values) == len(weights) && len(values) > 0 {
		wvalues := make([]float64, len(values))
		fitness := &Fitness{weights: weights, wvalues: wvalues, values: values, valid: true}
		fitness.SetValues(values)
		return fitness
	}
	return nil
}

// GetWValues returns the copy of wvalues
func (fitness *Fitness) GetWValues() []float64 {
	if fitness.weights == nil {
		return nil
	}
	wvalues := make([]float64, len(fitness.wvalues))
	copy(wvalues, fitness.wvalues)
	return wvalues
}

// GetWeights is used to get a copy of fitness's weights
func (fitness *Fitness) GetWeights() []float64 {
	weights := make([]float64, len(fitness.weights))
	copy(weights, fitness.weights)
	return weights
}

// GetValues is used to get the values of fitness
func (fitness *Fitness) GetValues() []float64 {
	values := make([]float64, len(fitness.weights))
	for i := range values {
		values[i] = fitness.wvalues[i] / fitness.weights[i]
	}
	return values
}

// SetValues is used to set the values of fitness
func (fitness *Fitness) SetValues(values []float64) {
	if fitness.wvalues == nil {
		fitness.wvalues = make([]float64, len(fitness.weights))
		fitness.values = make([]float64, len(fitness.weights))
	}
	for i, value := range values {
		fitness.wvalues[i] = value * fitness.weights[i]
		fitness.values[i] = value
	}
	fitness.valid = true
}

// Dominates is used to check if the fitness dominates the other
func (fitness *Fitness) Dominates(other *Fitness, obj []int) bool {
	notEqual := false
	if obj == nil {
		for i, wv := range fitness.wvalues {
			if wv > other.wvalues[i] {
				notEqual = true
			} else if wv < other.wvalues[i] {
				return false
			}
		}
	} else {
		for _, i := range obj {
			if fitness.wvalues[i] > other.wvalues[i] {
				notEqual = true
			} else if fitness.wvalues[i] < other.wvalues[i] {
				return false
			}
		}
	}
	return notEqual
}

// Clone returns a copy of fitness
func (fitness *Fitness) Clone() *Fitness {
	weights := make([]float64, len(fitness.weights))
	copy(weights, fitness.weights)
	if fitness.values != nil {
		values := make([]float64, len(fitness.values))
		copy(values, fitness.values)
		return NewFitnessWithValues(weights, values)
	}
	return NewFitness(weights)
}

// Valid is used to assess if a fitness is valid or not.
func (fitness *Fitness) Valid() bool {
	return fitness.valid && fitness.wvalues != nil && len(fitness.wvalues) != 0
}

// Invalidate makes the fitness value be invalid.
func (fitness *Fitness) Invalidate() {
	fitness.valid = false
}

// Greater is used to check if the fitness is greater than the other
func (fitness *Fitness) Greater(other *Fitness) bool {
	return !fitness.LessEqual(other)
}

// GreaterEqual is used to check if the fitness is greater than or equal to the other
func (fitness *Fitness) GreaterEqual(other *Fitness) bool {
	return !fitness.Less(other)
}

// LessEqual is used to check if the fitness is lessthan or equal to the other
func (fitness *Fitness) LessEqual(other *Fitness) bool {
	for i, fit := range fitness.wvalues {
		if other.wvalues[i] < fit {
			return false
		}
	}
	return true
}

// Less is used to check if the fitness is less than the other
func (fitness *Fitness) Less(other *Fitness) bool {
	for i, fit := range fitness.wvalues {
		if other.wvalues[i] <= fit {
			return false
		}
	}
	return true
}

// Equal is used to check if the fitness is euqal to the other
func (fitness *Fitness) Equal(other *Fitness) bool {
	for i, fit := range fitness.wvalues {
		if other.wvalues[i] != fit {
			return false
		}
	}
	return true
}

// NotEqual is used to check if the fitness is not euqal to the other
func (fitness *Fitness) NotEqual(other *Fitness) bool {
	return !fitness.Equal(other)
}

func (fitness *Fitness) String() string {
	fmtStr := "Fitness{weights:%v, values:%v, wvalues:%v, valid:%v}"
	return fmt.Sprintf(fmtStr, fitness.weights, fitness.values, fitness.wvalues, fitness.valid)
}

// Len returns the amounts of objective
func (fitness *Fitness) Len() int {
	return len(fitness.values)
}
