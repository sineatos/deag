package emo

import (
	"math"
	"sort"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/utility"
	"github.com/sineatos/deag/utility/list"
)

/*************************************
 * Non-Dominated Sorting   (NSGA-II) *
 *************************************/

// SelNSGA2 applies NSGA-II selection operator on the individuals.
// Usually, the size of individuals will be larger than k because any individual present in individuals will appear in the returned list at most once.
// Having the size of individuals equals to k will have no effect other than sorting the population according to their front rank.
// The list returned contains references to the input individuals.
// For more details on the NSGA-II operator see [Deb2002]_.
//
// individuals: A list of individuals to select from.
//
// k: The number of individuals to select.
//
// nd: Specify the non-dominated algorithm to use: 'standard' or 'log'.
//
// return: A list of selected individuals.
//
// [Deb2002] Deb, Pratab, Agarwal, and Meyarivan, "A fast elitist non-dominated sorting genetic algorithm for multi-objective optimization: NSGA-II", 2002.
func SelNSGA2(individuals base.Individuals, k int) base.Individuals {
	paretoFronts := sortNondominated(individuals, k, false)
	crowdingDist := make([][]float64, 0, len(paretoFronts))
	for _, front := range paretoFronts {
		cd := assignCrowdingDist(front)
		crowdingDist = append(crowdingDist, cd)
	}
	paretoFrontSize := len(paretoFronts)
	chosenSize := 0
	for _, front := range paretoFronts[:paretoFrontSize-1] {
		chosenSize += len(front)
	}
	var chosen base.Individuals

	k = k - chosenSize
	if k > 0 {
		chosen = make(base.Individuals, 0, k+chosenSize)
	} else {
		chosen = make(base.Individuals, 0, chosenSize)
	}

	for _, front := range paretoFronts[:paretoFrontSize-1] {
		chosen = append(chosen, front...)
	}

	if k > 0 {
		sorter := &crowdingDistIndividualSorter{
			Inds:         paretoFronts[paretoFrontSize-1],
			CrowdingDist: crowdingDist[paretoFrontSize-1],
		}
		sort.Sort(sort.Reverse(sorter))
		chosen = append(chosen, sorter.Inds[:k]...)
	}
	return chosen
}

// sortNondominated sorts the first k individuals into different nondomination levels using the "Fast Nondominated Sorting Approach" proposed by Deb et al., see [Deb2002]_. This algorithm has a time complexity of `O(MN^2)`, where `M` is the number of objectives and `N` the number of individuals.
//
// individuals: A list of individuals to select from.
//
// k: The number of individuals to select.
//
// firstFrontOnly: If true sort only the first front and exit.
//
// returns a slice of Pareto fronts (base.Individuals), the first list includes nondominated individuals.
//
// [Deb2002] Deb, Pratab, Agarwal, and Meyarivan, "A fast elitist non-dominated sorting genetic algorithm for multi-objective optimization: NSGA-II", 2002.
func sortNondominated(individuals base.Individuals, k int, firstFrontOnly bool) []base.Individuals {
	if k == 0 {
		return []base.Individuals{}
	}

	fits := uniqueByFitness(individuals)

	cap := len(individuals)
	currentFront := make([]*base.Fitness, 0, cap)
	nextFront := make([]*base.Fitness, 0, cap)
	dominatingFits := make(map[*base.Fitness]int)
	dominatedFits := list.NewStaticLinkedList(cap * (cap - 1) / 2)

	for i, fitI := range fits {
		for j, fitJ := range fits[i+1:] {
			if fitI.Dominates(fitJ, nil) {
				dominatingFits[fitJ]++
				dominatedFits.Add(fitI, j)
			} else if fitJ.Dominates(fitI, nil) {
				dominatingFits[fitI]++
				dominatedFits.Add(fitJ, i)
			}
		}
		if dominatingFits[fitI] == 0 {
			currentFront = append(currentFront, fitI)
		}
	}

	mapFitInd := mapIndividualsByFitness(fits, individuals)
	fronts, frontIndex := list.NewStaticLinkedList(cap), 0

	extends := func(fit *base.Fitness) {
		next, exist := mapFitInd.GetFirstDataIndex(fit)
		var data int
		if exist {
			for next != -1 {
				data, next = mapFitInd.GetData(next)
				fronts.Add(frontIndex, data)
			}
		}
	}

	for _, fit := range currentFront {
		extends(fit)
	}
	paretoSorted, _ := fronts.GetSize(frontIndex)

	// Rank the next front until all individuals are sorted or the given number of individual are sorted.
	if !firstFrontOnly {
		N := utility.If(cap < k, cap, k).(int)
		frontIndex++
		for paretoSorted < N {
			for _, fitP := range currentFront {
				next, exist := dominatedFits.GetFirstDataIndex(fitP)
				if exist {
					fitDIdx, next := dominatedFits.GetData(next)
					for next != -1 {
						fitD := fits[fitDIdx]
						dominatingFits[fitD]--
						if dominatingFits[fitD] == 0 {
							nextFront = append(nextFront, fitD)
							fSize, _ := mapFitInd.GetSize(fitD)
							paretoSorted += fSize
							extends(fitD)
						}
						fitDIdx, next = dominatedFits.GetData(next)
					}
				}
			}
			currentFront, nextFront = nextFront, currentFront
			nextFront = nextFront[0:0]
		}
	}

	frontSlice := staticLinkedList2Slice(fronts, individuals)
	return frontSlice
}

// uniqueByFitness sorts the individuals and remove all individuals having the same fitness value and keeps individuals which fitness value is different, the input individuals do not be changed, finally returns the unique individuals
func uniqueByFitness(individuals base.Individuals) []*base.Fitness {
	fits := make([]*base.Fitness, len(individuals))
	for i, ind := range individuals {
		fits[i] = ind.GetFitness()
	}
	sorter := &fitnessSorter{FitList: fits}
	sort.Sort(sorter)
	current := 0
	for i := 1; i < len(fits); i++ {
		if !fits[current].Equal(fits[i]) {
			current++
			fits[current] = fits[i]
		}
	}
	return fits[:current+1]
}

// mapIndividualsByFitness classifies all individuals accroding to their fitness, and returns a static linked list having the information
func mapIndividualsByFitness(uniqueFitness []*base.Fitness, individuals base.Individuals) *list.StaticLinkedList {
	mapping := list.NewStaticLinkedList(len(individuals))
	size := len(uniqueFitness)
	for i, ind := range individuals {
		indFit := ind.GetFitness()
		indWValues := indFit.GetWValues()
		loc := sort.Search(size, func(j int) bool {
			for k, wv := range uniqueFitness[j].GetWValues() {
				if wv != indWValues[k] {
					return wv > indWValues[k]
				}
			}
			return true
		})
		if loc < size && uniqueFitness[loc].Equal(indFit) {
			mapping.Add(uniqueFitness[loc], i)
		} else {
			panic("It's impossible that cannot find a fitness in uniqueFitness that equals to ind's fitness")
		}
	}
	return mapping
}

// staticLinkedList returns a slice of base.Individuals, each element of slice is a base.Individuals which is the pareto front
func staticLinkedList2Slice(fronts *list.StaticLinkedList, individuals base.Individuals) []base.Individuals {
	head := fronts.GetHead()
	headSize := len(head)
	ans := make([]base.Individuals, headSize)
	for i := 0; i < headSize; i++ {
		next, exist := fronts.GetFirstDataIndex(i)
		var idx int
		if exist {
			listSize, _ := fronts.GetSize(i)
			ans[i] = make(base.Individuals, 0, listSize)
			for next != -1 {
				idx, next = fronts.GetData(next)
				ans[i] = append(ans[i], individuals[idx])
			}
		}
	}
	return ans
}

// assignCrowdingDist assigns a crowding distance to each individual's fitness, and returns a slice of float64 which is the crowding distance of individual's fitness.
func assignCrowdingDist(individuals base.Individuals) []float64 {
	popSize := len(individuals)
	if popSize == 0 {
		return []float64{}
	}
	distances := make([]float64, popSize)
	indDist := make([]float64, 0, popSize)

	crowd := &crowdingDistSorter{
		ValuesSlice: make([][]float64, 0, popSize),
		I:           make([]int, 0, popSize),
		Loc:         0,
	}
	for i, ind := range individuals {
		crowd.ValuesSlice = append(crowd.ValuesSlice, ind.GetFitness().GetValues())
		crowd.I = append(crowd.I, i)
	}

	nobJ := individuals[0].GetFitness().Len()
	for i := 0; i < nobJ; i++ {
		crowd.Loc = i
		sort.Sort(crowd)
		distances[crowd.I[0]], distances[crowd.I[popSize-1]] = math.Inf(1), math.Inf(1)
		diff := crowd.ValuesSlice[popSize-1][i] - crowd.ValuesSlice[0][i]
		if diff == 0.0 {
			continue
		}
		norm := float64(nobJ) * (diff)
		var prev, cur, next int
		for prev = range crowd.ValuesSlice[:popSize-2] {
			cur, next = prev+1, prev+2
			distances[crowd.I[cur]] += (crowd.ValuesSlice[next][i] - crowd.ValuesSlice[prev][i]) / norm
		}
	}

	for _, dist := range distances {
		indDist = append(indDist, dist)
	}

	return indDist
}

// SelTournamentDCD is tournament selection based on dominance (D) between two individuals, if the two individuals do not interdominate the selection is made based on crowding distance (CD).
// The individuals sequence length has to be a multiple of 4.
// Starting from the beginning of the selected individuals, two consecutive individuals will be different (assuming all individuals in the input list are unique).
// Each individual from the input list won't be selected more than twice.
//
// This selection requires the individuals to have a :attr:`crowding_dist` attribute, which can be set by the :func:`assignCrowdingDist` function.
//
// individuals: A list of individuals to select from.
//
// k: The number of individuals to select.
//
// returns: A list of selected individuals.
// func SelTournamentDCD(individuals base.Individuals, k int) base.Individuals {
// 	return nil
// }

type crowdingDistSorter struct {
	ValuesSlice [][]float64
	I           []int
	Loc         int
}

func (s *crowdingDistSorter) Len() int {
	return len(s.I)
}

func (s *crowdingDistSorter) Less(i, j int) bool {
	return s.ValuesSlice[i][s.Loc] < s.ValuesSlice[j][s.Loc]
}

func (s *crowdingDistSorter) Swap(i, j int) {
	s.ValuesSlice[i], s.ValuesSlice[j] = s.ValuesSlice[j], s.ValuesSlice[i]
	s.I[i], s.I[j] = s.I[j], s.I[i]
}

type crowdingDistIndividualSorter struct {
	Inds         base.Individuals
	CrowdingDist []float64
}

func (s *crowdingDistIndividualSorter) Len() int {
	return len(s.Inds)
}

func (s *crowdingDistIndividualSorter) Less(i, j int) bool {
	return s.CrowdingDist[i] < s.CrowdingDist[j]
}

func (s *crowdingDistIndividualSorter) Swap(i, j int) {
	s.Inds[i], s.Inds[j] = s.Inds[j], s.Inds[i]
	s.CrowdingDist[i], s.CrowdingDist[j] = s.CrowdingDist[j], s.CrowdingDist[i]
}

type fitnessSorter struct {
	FitList []*base.Fitness
}

func (s *fitnessSorter) Len() int {
	return len(s.FitList)
}

func (s *fitnessSorter) Less(i, j int) bool {
	valI := s.FitList[i].GetWValues()
	valJ := s.FitList[j].GetWValues()
	for k, vi := range valI {
		if vi != valJ[k] {
			return vi < valJ[k]
		}
	}
	return true
}

func (s *fitnessSorter) Swap(i, j int) {
	s.FitList[i], s.FitList[j] = s.FitList[j], s.FitList[i]
}
