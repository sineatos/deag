package support

import (
	"github.com/sineatos/deag/base"
)

// ParetoFront is pareto front hall of fame contains all the non-dominated individuals that ever lived in the population. That means that the Pareto front hall of fame can contain an infinity of different individuals.
type ParetoFront HallOfFame

// DefaultParetoFront is a default ParetoFront implement
type DefaultParetoFront struct {
	DefaultHallOfFame
}

// NewDefaultParetoFront returns a DefaultParetoFront setting with similar.
// If similar is nil, sets the similar as
//  func(x, y base.Individual) bool {x.IsEqual(y)}
func NewDefaultParetoFront(similar func(base.Individual, base.Individual) bool) *DefaultParetoFront {
	hof := NewDefaultHallOfFame(0, similar)
	return &DefaultParetoFront{DefaultHallOfFame: *hof}
}

// Update the Pareto front hall of fame with the individuals by adding the individuals from the population that are not dominated by the hall of fame.
// If any individual in the hall of fame is dominated it is removed.
func (pf *DefaultParetoFront) Update(individuals base.Individuals) {
	toRemove, count := make([]int, 0, len(individuals)), 0
	for _, ind := range individuals {
		isDominated, dominatesOne, hasTwin := false, false, false
		count = 0
		indFitness := ind.GetFitness()
		for i, hofer := range pf.items {
			hofFitness := hofer.GetFitness()
			if !dominatesOne && hofFitness.Dominates(indFitness, nil) {
				isDominated = true
				break
			} else if indFitness.Dominates(hofFitness, nil) {
				dominatesOne = true
				toRemove[count] = i
				count++
			} else if indFitness.Equal(hofFitness) && pf.similar(ind, hofer) {
				hasTwin = true
				break
			}
		}
		for i := count - 1; i >= 0; i-- {
			pf.Remove(toRemove[i])
		}
		if !isDominated && !hasTwin {
			pf.Insert(ind)
		}
	}
}
