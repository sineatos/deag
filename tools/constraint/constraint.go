package constraint

import (
	"fmt"
	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
)

// Float64Constraint defines the functions of constraint
type Float64Constraint interface {
	// AdjustAndEvolve checks individual and adjusts it if it is not feasible
	AdjustAndEvolve(individual *base.Float64Individual) []float64
}

// ClosestValidPenalty returns penalized fitness for invalid individuals and the original fitness value for valid individuals. The penalized fitness is made of the fitness of the closest valid individual added with a weighted (optional) distance penalty. The distance function, if provided, shall return a value growing as the individual moves away the valid zone.
type ClosestValidPenalty struct {
	// isFeasible checks individual if feasible
	isFeasible func(individual *base.Float64Individual) bool
	// adjust adjusts individual and returns the individual in constraint
	adjust    func(individual *base.Float64Individual) *base.Float64Individual
	alpha     float64
	distance  func(ind1, ind2 *base.Float64Individual) float64
	evaluator benchmarks.Float64Evaluator
}

// NewClosestValidPenalty returns a *ClosestValidPenalty
func NewClosestValidPenalty(isFeasible func(individual *base.Float64Individual) bool, adjust func(individual *base.Float64Individual) *base.Float64Individual, alpha float64, distance func(ind1, ind2 *base.Float64Individual) float64, evaluator benchmarks.Float64Evaluator) *ClosestValidPenalty {
	return &ClosestValidPenalty{
		isFeasible: isFeasible,
		adjust:     adjust,
		alpha:      alpha,
		distance:   distance,
		evaluator:  evaluator,
	}
}

// AdjustAndEvolve checks individual and adjusts it if it is not feasible
func (con *ClosestValidPenalty) AdjustAndEvolve(individual *base.Float64Individual) []float64 {
	if con.isFeasible(individual) {
		return con.evaluator(individual)
	}
	fInd := con.adjust(individual)
	fAns := con.evaluator(fInd)
	fWeights := individual.GetFitness().GetWeights()
	weights := make([]float64, len(fWeights))
	for i := range fWeights {
		if fWeights[i] >= 0.0 {
			weights[i] = 1.0
		} else {
			weights[i] = -1.0
		}
	}
	if len(weights) != len(fAns) {
		panic(fmt.Sprintf("Fitness weights and computed fitness are of different size: %v,%v", weights, fAns))
	}
	dists := make([]float64, len(weights))
	if con.distance != nil {
		d := con.distance(fInd, individual)
		for i := range dists {
			dists[i] = d
		}
	}
	ans := make([]float64, len(dists))
	for i := range ans {
		ans[i] = fAns[i] - weights[i]*con.alpha*dists[i]
	}
	return ans
}
