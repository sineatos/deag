package crossover

import (
	"math/rand"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/utility"
)

/***************************
 * GA Crossovers           *
 ***************************/

// CxOnePointInt executes a one point crossover on the input individuals.
//
// The two individuals are modified in place. The resulting individuals will respectively have the length of the other.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// Returns:
//
// The crossovered individuals.
func CxOnePointInt(ind1 *base.IntIndividual, ind2 *base.IntIndividual) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	cxpoint := 1 + rand.Intn(size-1)
	c1, c2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	if l1 <= l2 {
		temp := make([]int, l2)
		copy(temp, c1[:cxpoint])
		copy(temp[cxpoint:], c2[cxpoint:])
		copy(c2[cxpoint:], c1[cxpoint:])
		c1 = temp
		c2 = c2[:l1]
	} else {
		temp := make([]int, l1)
		copy(temp, c2[:cxpoint])
		copy(temp[cxpoint:], c1[cxpoint:])
		copy(c1[cxpoint:], c2[cxpoint:])
		c2 = temp
		c1 = c1[:l2]
	}
	ind1.SetChromosome(c1)
	ind2.SetChromosome(c2)
	return ind1, ind2
}

// CxTwoPointInt Executes a two-point crossover on the input individuals.
// The two individuals are modified in place and both keep their original length.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// Returns:
//
// The crossovered individuals.
func CxTwoPointInt(ind1 *base.IntIndividual, ind2 *base.IntIndividual) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	cxpoint1 := 1 + rand.Intn(size)
	cxpoint2 := 1 + rand.Intn(size-1)
	if cxpoint2 >= cxpoint1 {
		cxpoint2++
	} else { // Swap the two cx points
		cxpoint1, cxpoint2 = cxpoint2, cxpoint1
	}
	c1, c2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	var temp int
	for i := cxpoint1; i < cxpoint2; i++ {
		temp = c1[i]
		c1[i] = c2[i]
		c2[i] = temp
	}
	return ind1, ind2
}

// CxUniformInt executes a uniform crossover that modify in place the two individuals.
// The attributes are swapped according to the indpb probability
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// indpb(float64): Independent probabily for each attribute to be exchanged.
//
// Returns:
//
// The crossovered individuals.
func CxUniformInt(ind1 *base.IntIndividual, ind2 *base.IntIndividual, indpb float64) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	c1, c2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			c1[i], c2[i] = c2[i], c1[i]
		}
	}
	return ind1, ind2
}

// CxPartialyMatched executes a partially matched crossover (PMX) on the input individuals.
// The two individuals are modified in place.
// This crossover expects individuals of indices(integer), the result for any other type of individuals is unpredictable.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// Returns:
//
// The crossovered individuals.
//
// Moreover, this crossover generates two children by matching pairs of values in a certain range of the two parents and swapping the values of those indexes. For more details see [Goldberg1985].
//
// [Goldberg1985] Goldberg and Lingel, "Alleles, loci, and the traveling salesman problem", 1985.
func CxPartialyMatched(ind1 *base.IntIndividual, ind2 *base.IntIndividual) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	p1, p2 := make([]int, size), make([]int, size)
	c1, c2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	// Initialize the position of each indices in the individuals
	for i := 0; i < size; i++ {
		p1[c1[i]], p2[c2[i]] = i, i
	}
	// Choose crossover points
	cxpoint1 := rand.Intn(size + 1)
	cxpoint2 := rand.Intn(size)
	if cxpoint2 >= cxpoint1 {
		cxpoint2++
	} else { // Swap the two cx points
		cxpoint1, cxpoint2 = cxpoint2, cxpoint1
	}
	// Apply crossover between cx points
	for i := cxpoint1; i < cxpoint2; i++ {
		// Keep track of the selected values
		temp1, temp2 := c1[i], c2[i]
		// Swap the matched value
		c1[i], c1[p1[temp2]] = temp2, temp1
		c2[i], c1[p2[temp1]] = temp1, temp2
		// Position bookkeeping
		p1[temp1], p1[temp2] = p1[temp2], p1[temp1]
		p2[temp1], p2[temp2] = p2[temp2], p2[temp1]
	}
	return ind1, ind2
}

// CxUniformPartialyMatched executes a uniform partially matched crossover (UPMX) on the input individuals.
// The two individuals are modified in place.
// This crossover expects individuals of indices(integer), the result for any other type of individuals is unpredictable.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// indpb(float64): Independent probabily for each attribute to be exchanged.
//
// Returns:
//
// The crossovered individuals.
//
// Moreover, this crossover generates two children by matching pairs of values chosen at random with a probability of indpb in the two parents and swapping the values of those indexes.
// For more details see [Cicirello2000].
//
// [Cicirello2000] Cicirello and Smith, "Modeling GA performance for control parameter optimization", 2000.
func CxUniformPartialyMatched(ind1 *base.IntIndividual, ind2 *base.IntIndividual, indpb float64) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	p1, p2 := make([]int, size), make([]int, size)
	c1, c2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	for i := 0; i < size; i++ {
		p1[c1[i]], p2[c2[i]] = i, i
	}
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			// Keep track of the selected values
			temp1, temp2 := c1[i], c2[i]
			// Swap the matched value
			c1[i], c1[p1[temp2]] = temp2, temp1
			c2[i], c1[p2[temp1]] = temp1, temp2
			// Position bookkeeping
			p1[temp1], p1[temp2] = p1[temp2], p1[temp1]
			p2[temp1], p2[temp2] = p2[temp2], p2[temp1]
		}
	}
	return ind1, ind2
}

// CxOrdered executes an ordered crossover (OX) on the input individuals.
// The two individuals are modified in place.
// his crossover expects individuals of indices, the result for any other type of individuals is unpredictable.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// Returns:
//
// The crossovered individuals.
//
// Moreover, this crossover generates holes in the input individuals.
// A hole is created when an attribute of an individual is between the two crossover points of the other individual.
// Then it rotates the element so that all holes are between the crossover points and fills them with the removed elements in order.
// For more details see [Goldberg1989].
//
// [Goldberg1989] Goldberg. Genetic algorithms in search, optimization and machine learning. Addison Wesley, 1989
func CxOrdered(ind1 *base.IntIndividual, ind2 *base.IntIndividual) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	a, b := rand.Intn(size), rand.Intn(size)
	c1, c2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	holes1, holes2 := make([]bool, size), make([]bool, size)
	if a > b {
		a, b = b, a
	}
	for i := 0; i < size; i++ {
		if i < a || i > b {
			holes1[i], holes2[i] = false, false
		} else {
			holes1[i], holes2[i] = true, true
		}
	}
	// We must keep the original values somewhere before scrambling everything
	temp1, temp2 := make([]int, len(c1)), make([]int, len(c2))
	for i, g := range c1 {
		temp1[i] = g
	}
	for i, g := range c2 {
		temp2[i] = g
	}
	k1, k2 := b+1, b+1
	for i := 0; i < size; i++ {
		if !holes1[temp1[(i+b+1)%size]] {
			c1[k1%size] = temp1[(i+b+1)%size]
			k1++
		}
		if !holes2[temp2[(i+b+1)%size]] {
			c2[k2%size] = temp2[(i+b+1)%size]
			k2++
		}
	}
	// Swap the content between a and b (included)
	for i := a; i < b+1; i++ {
		c1[i], c2[i] = c2[i], c1[i]
	}
	return ind1, ind2
}

/***************************
 * Messy Crossovers        *
 ***************************/

// CxMessyOnePointInt executes a one point crossover on individual.
// The crossover will in most cases change the individuals size.
// The two individuals are modified in place.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// Returns:
//
// The crossovered individuals.
func CxMessyOnePointInt(ind1 *base.IntIndividual, ind2 *base.IntIndividual) (*base.IntIndividual, *base.IntIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	cxpoint1, cxpoint2 := rand.Intn(l1+1), rand.Intn(l2+1)
	r1, r2 := l1-cxpoint1, l2-cxpoint2
	gc1, gc2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	temp1, temp2 := make([]int, cxpoint1+r2), make([]int, cxpoint2+r1)
	copy(temp1, gc1[:cxpoint1])
	copy(temp1[cxpoint1:], gc2[cxpoint2:])
	copy(temp2, gc2[:cxpoint2])
	copy(temp2, gc1[cxpoint1:])
	ind1.SetChromosome(gc1)
	ind2.SetChromosome(gc2)
	return ind1, ind2
}

/***************************
 * ES Crossovers           *
 ***************************/

// CxESTwoPointInt executes a classical two points crossover on both the individuals and their strategy.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// Returns:
//
// The crossovered individuals.
func CxESTwoPointInt(ind1 *base.IntESIndividual, ind2 *base.IntESIndividual, alpha float64) (*base.IntESIndividual, *base.IntESIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	gc1, gc2 := ind1.GetChromosome().([]int), ind2.GetChromosome().([]int)
	st1, st2 := ind1.GetStrategies(), ind2.GetStrategies()
	pt1, pt2 := 1+rand.Intn(size), 1+rand.Intn(size-1)
	if pt2 >= pt1 {
		pt2++
	} else { // Swap the two cx points
		pt1, pt2 = pt2, pt1
	}
	for i := pt1; i < pt2; i++ {
		tmp1, tmp2 := gc1[i], st1[i]
		gc1[i], st1[i] = gc2[i], st2[i]
		gc2[i], st2[i] = tmp1, tmp2
	}
	return ind1, ind2
}
