package crossover

import (
	"math"
	"math/rand"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/utility"
)

/***************************
 * GA Crossovers           *
 ***************************/

// CxOnePointFloat64 executes a one point crossover on the input individuals.
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
func CxOnePointFloat64(ind1 *base.Float64Individual, ind2 *base.Float64Individual) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	cxpoint := 1 + rand.Intn(size-1)
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	if l1 <= l2 {
		temp := make([]float64, l2)
		copy(temp, gc1[:cxpoint])
		copy(temp[cxpoint:], gc2[cxpoint:])
		copy(gc2[cxpoint:], gc1[cxpoint:])
		gc1 = temp
		gc2 = gc2[:l1]
	} else {
		temp := make([]float64, l1)
		copy(temp, gc2[:cxpoint])
		copy(temp[cxpoint:], gc1[cxpoint:])
		copy(gc1[cxpoint:], gc2[cxpoint:])
		gc2 = temp
		gc1 = gc1[:l2]
	}
	ind1.SetChromosome(gc1)
	ind2.SetChromosome(gc2)
	return ind1, ind2
}

// CxTwoPointFloat64 Executes a two-point crossover on the input individuals.
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
func CxTwoPointFloat64(ind1 *base.Float64Individual, ind2 *base.Float64Individual) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	cxpoint1 := 1 + rand.Intn(size)
	cxpoint2 := 1 + rand.Intn(size-1)
	if cxpoint2 >= cxpoint1 {
		cxpoint2++
	} else { // Swap the two cx points
		cxpoint1, cxpoint2 = cxpoint2, cxpoint1
	}
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	var temp float64
	for i := cxpoint1; i < cxpoint2; i++ {
		temp = gc1[i]
		gc1[i] = gc2[i]
		gc2[i] = temp
	}
	return ind1, ind2
}

// CxUniformFloat64 executes a uniform crossover that modify in place the two individuals.
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
func CxUniformFloat64(ind1 *base.Float64Individual, ind2 *base.Float64Individual, indpb float64) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		if rand.Float64() < indpb {
			gc1[i], gc2[i] = gc2[i], gc1[i]
		}
	}
	return ind1, ind2
}

// CxBlend executes a blend crossover that modify in-place the input individuals.
// The blend crossover expects individuals of floating point numbers.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// alpha(float64): Extent of the interval in which the new values can be drawn for each attribute on both side of the parents' attributes.
//
// Returns:
//
// The crossovered individuals.
func CxBlend(ind1 *base.Float64Individual, ind2 *base.Float64Individual, alpha float64) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		gamma := (1.+2.*alpha)*rand.Float64() - alpha
		x1, x2 := gc1[i], gc2[i]
		gc1[i] = (1.-gamma)*x1 + gamma*x2
		gc2[i] = gamma*x1 + (1.-gamma)*x2
	}
	return ind1, ind2
}

// CxSimulatedBinary executes a simulated binary crossover that modify in-place the input individuals.
// The simulated binary crossover expects individuals of floating point numbers.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// eta(float64): Crowding degree of the crossover.
// A high eta will produce children resembling to their parents, while a small eta will produce solutions much more different.
//
// Returns:
//
// The crossovered individuals.
func CxSimulatedBinary(ind1 *base.Float64Individual, ind2 *base.Float64Individual, eta float64) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		ran := rand.Float64()
		var beta float64
		if ran <= 0.5 {
			beta = 2. * ran
		} else {
			beta = 1. / (2. * (1. - ran))
		}
		beta = math.Pow(beta, 1./(eta+1.))
		x1, x2 := gc1[i], gc2[i]
		gc1[i] = 0.5 * (((1 + beta) * x1) + ((1 - beta) * x2))
		gc2[i] = 0.5 * (((1 - beta) * x1) + ((1 + beta) * x2))
	}
	return ind1, ind2
}

// CxSimulatedBinaryBounded Executes a simulated binary crossover that modify in-place the input individuals.
// The simulated binary crossover expects individuals of floating point numbers.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// eta(float64): Crowding degree of the crossover.
// A high eta will produce children resembling to their parents, while a small eta will produce solutions much more different.
//
// low(float or []float64): values that is the lower bound of the search space.
//
// up(float or []float64): values that is the lower bound of the search space.
//
// Returns:
//
// The crossovered individuals.
//
// Note: This implementation is similar to the one implemented in the original NSGA-II C code presented by Deb.
func CxSimulatedBinaryBounded(ind1 *base.Float64Individual, ind2 *base.Float64Individual, eta float64, low interface{}, up interface{}) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	lows := utility.Interface2Float64Slice("low", low, size)
	ups := utility.Interface2Float64Slice("up", up, size)
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	for i := 0; i < size; i++ {
		y1, y2 := gc1[i], gc2[i]
		if rand.Float64() <= 0.5 {
			if math.Abs(y1-y2) > 1E-14 {
				x1 := math.Min(y1, y2)
				x2 := math.Max(y1, y2)
				ran := rand.Float64()

				beta := 1.0 + (2.0 * (x1 - lows[i]) / (x2 - x1))
				alpha := 2.0 - math.Pow(beta, -(eta+1))
				var betaQ float64

				if ran <= 1.0/alpha {
					betaQ = math.Pow(ran*alpha, 1.0/(eta+1))
				} else {
					betaQ = math.Pow(1.0/(2.0-ran*alpha), 1.0/(eta+1))
				}
				c1 := 0.5 * (x1 + x2 - betaQ*(x2-x1))

				beta = 1.0 + (2.0 * (ups[i] - x2) / (x2 - x1))
				alpha = 2.0 - math.Pow(beta, -(eta+1))
				if ran <= 1.0/alpha {
					betaQ = math.Pow(ran*alpha, 1.0/(eta+1))
				} else {
					betaQ = math.Pow(1.0/(2.0-ran*alpha), 1.0/(eta+1))
				}
				c2 := 0.5 * (x1 + x2 + betaQ*(x2-x1))

				c1 = math.Min(math.Max(c1, lows[i]), ups[i])
				c2 = math.Min(math.Max(c2, lows[i]), ups[i])

				if rand.Float64() <= 0.5 {
					gc1[i] = c2
					gc2[i] = c1
				} else {
					gc1[i] = c1
					gc2[i] = c2
				}
			}
		}
	}
	return ind1, ind2
}

/***************************
 * Messy Crossovers        *
 ***************************/

// CxMessyOnePointFloat64 executes a one point crossover on individual.
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
func CxMessyOnePointFloat64(ind1 *base.Float64Individual, ind2 *base.Float64Individual) (*base.Float64Individual, *base.Float64Individual) {
	l1, l2 := ind1.Len(), ind2.Len()
	cxpoint1, cxpoint2 := rand.Intn(l1+1), rand.Intn(l2+1)
	r1, r2 := l1-cxpoint1, l2-cxpoint2
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	temp1, temp2 := make([]float64, cxpoint1+r2), make([]float64, cxpoint2+r1)
	copy(temp1, gc1[:cxpoint1])
	copy(temp1[cxpoint1:], gc2[cxpoint2:])
	copy(temp2, gc2[:cxpoint2])
	copy(temp2[cxpoint2:], gc1[cxpoint1:])
	ind1.SetChromosome(temp1)
	ind2.SetChromosome(temp2)
	return ind1, ind2
}

/***************************
 * ES Crossovers           *
 ***************************/

// CxESBlend executes a blend crossover on both, the individual and the strategy.
// Adjustement of the minimal strategy shall be done after the call to this function, consider using a decorator.
//
// Parameters:
//
// ind1: The first individual participating in the crossover.
//
// ind2: The second individual participating in the crossover.
//
// alpha(float64): Extent of the interval in which the new values can be drawn for each attribute on both side of the parents' attributes.
//
// Returns:
//
// The crossovered individuals.
func CxESBlend(ind1 *base.Float64ESIndividual, ind2 *base.Float64ESIndividual, alpha float64) (*base.Float64ESIndividual, *base.Float64ESIndividual) {
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
	st1, st2 := ind1.GetStrategies(), ind2.GetStrategies()
	gc1l, gc2l, st1l, st2l := len(gc1), len(gc2), len(st1), len(st2)
	gsl := utility.If(gc1l > gc2l, gc1l, gc2l).(int)
	stl := utility.If(st1l > st2l, st1l, st2l).(int)
	size := utility.If(gsl > stl, gsl, stl).(int)
	for i := 0; i < size; i++ {
		// Blend the values
		x1, x2, s1, s2 := gc1[i], gc2[i], st1[i], st2[i]
		gamma := (1.+2.*alpha)*rand.Float64() - alpha
		gc1[i] = (1.-gamma)*x1 + gamma*x2
		gc2[i] = gamma*x1 + (1.-gamma)*x2
		// Blend the strategies
		gamma = (1.+2.*alpha)*rand.Float64() - alpha
		st1[i] = (1.-gamma)*s1 + gamma*s2
		st2[i] = gamma*s1 + (1.-gamma)*s2

	}
	return ind1, ind2
}

// CxESTwoPointFloat64 executes a classical two points crossover on both the individuals and their strategy.
// The individuals shall be  base.ESIndividual.
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
func CxESTwoPointFloat64(ind1 *base.Float64ESIndividual, ind2 *base.Float64ESIndividual, alpha float64) (*base.Float64ESIndividual, *base.Float64ESIndividual) {
	l1, l2 := ind1.Len(), ind2.Len()
	size := utility.If(l1 < l2, l1, l2).(int)
	gc1, gc2 := ind1.GetChromosome().([]float64), ind2.GetChromosome().([]float64)
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
