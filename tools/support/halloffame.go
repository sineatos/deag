package support

import (
	"fmt"
	"sort"

	"github.com/sineatos/deag/base"
)

// HallOfFame contains the best individual that ever lived in the population during the evolution.
// It is lexicographically sorted at all time so that the first element of the hall of fame is the individual that has the best first fitness value ever seen,
// according to the weights provided to the fitness at creation time.
type HallOfFame interface {
	// Update the hall of fame with the individuals by replacing the worst individuals in it by the best individuals present in individuals (if they are better).
	Update(individuals base.Individuals)
	// Insert inserts a new individual in the hall of fame using the function.
	Insert(ind base.Individual)
	// Remove removes the specified index from the hall of fame.
	Remove(index int)
	// Clear clears the hall of fame.
	Clear()
	// Len returns the size of items
	Len() int
	// Get returns the Individual of index
	Get(index int) base.Individual
	// Reversed returns a copy of reversed HallOfFame
	Reversed() base.Individuals
	// String returns string of HallOfFame
	String() string
}

// DefaultHallOfFame is a default HallOfFame implement
type DefaultHallOfFame struct {
	maxsize int
	size    int
	items   base.Individuals // sorted items using Fitness from best to worse
	similar func(base.Individual, base.Individual) bool
}

// NewDefaultHallOfFame returns a DefaultHallOfFame setting with maxsize and similar.
// If similar is nil, sets the similar as
//	func(x, y base.Individual) bool { x.IsEqual(y) }
func NewDefaultHallOfFame(maxsize int, similar func(base.Individual, base.Individual) bool) *DefaultHallOfFame {
	if similar == nil {
		similar = func(x, y base.Individual) bool {
			return x.IsEqual(y)
		}
	}
	items := make([]base.Individual, maxsize)
	return &DefaultHallOfFame{maxsize: maxsize, size: 0, items: items, similar: similar}
}

// Update the hall of fame with the individuals by replacing the worst individuals in it by the best individuals present in individuals (if they are better).
// The size of the hall of fame is kept constant.
func (hof *DefaultHallOfFame) Update(individuals base.Individuals) {
	if hof.size == 0 && hof.maxsize != 0 && individuals.Len() > 0 {
		hof.Insert(individuals[0])
	}
	for _, ind := range individuals {
		if ind.GetFitness().Greater(hof.items[hof.size-1].GetFitness()) || hof.size < hof.maxsize {
			flag := true
			for _, hofer := range hof.items {
				if hof.similar(ind, hofer) {
					flag = false
					break
				}
			}
			if flag {
				// The individual is unique and strictly better than the worst
				if hof.size >= hof.maxsize {
					hof.Remove(hof.maxsize)
				}
				hof.Insert(ind)
			}
		}
	}
}

// Insert inserts a new individual in the hall of fame using the function.
// The inserted individual is inserted on the right side of an equal individual.
// Inserting a new individual in the hall of fame also preserve the hall of fame's order.
// This method does not check for the size of the hall of fame, in a way that inserting a new individual in a full hall of fame will not remove the worst individual to maintain a constant size.
func (hof *DefaultHallOfFame) Insert(ind base.Individual) {
	ind = ind.Clone().(base.Individual)
	indFitness := ind.GetFitness()
	index := sort.Search(hof.size, func(i int) bool {
		return hof.items[i].GetFitness().Less(indFitness)
	})
	if hof.size < hof.maxsize {
		hof.size++
	}
	for i := hof.size - 1; i > index; i-- {
		hof.items[i] = hof.items[i-1]
	}
	hof.items[index] = ind
}

// Remove removes the specified index from the hall of fame.
func (hof *DefaultHallOfFame) Remove(index int) {
	if hof.size != 0 {
		hof.size--
		for i := index; i < hof.size; i++ {
			hof.items[i] = hof.items[i+1]
		}
	}
}

// Clear clears the hall of fame.
func (hof *DefaultHallOfFame) Clear() {
	hof.size = 0
}

// Len returns the size of items
func (hof *DefaultHallOfFame) Len() int {
	return hof.size
}

// Get returns the Individual of index
func (hof *DefaultHallOfFame) Get(index int) base.Individual {
	return hof.items[index]
}

// Reversed returns a copy of reversed HallOfFame
func (hof *DefaultHallOfFame) Reversed() base.Individuals {
	ans := make(base.Individuals, hof.size)
	for i, j := hof.size-1, 0; i >= 0; i-- {
		ans[j] = hof.items[i]
		j++
	}
	return ans
}

// String returns string of HallOfFame
func (hof *DefaultHallOfFame) String() string {
	return fmt.Sprintf("%v", hof.items[:hof.size])
}
