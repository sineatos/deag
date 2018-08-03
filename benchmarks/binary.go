package benchmarks

import (
	"math"

	"github.com/sineatos/deag/base"
)

// Trap is binary benchmark
func Trap(individual *base.BoolIndividual) []float64 {
	u, k := 0, individual.Len()
	for _, val := range individual.GetChromosome().([]bool) {
		if val {
			u++
		}
	}
	if u == k {
		return []float64{float64(k)}
	}
	return []float64{float64(k - 1 - u)}
}

// InvTrap is binary benchmark
func InvTrap(individual *base.BoolIndividual) []float64 {
	u, k := 0, individual.Len()
	for _, val := range individual.GetChromosome().([]bool) {
		if val {
			u++
		}
	}
	if u == 0 {
		return []float64{float64(k)}
	}
	return []float64{float64(u - 1)}
}

// ChuangF1 is maximization binary benchmark
//
// Binary deceptive function from : Multivariate Multi-Model Approach for Globally Multimodal Problems by Chung-Yao Chuang and Wen-Lian Hsu.
//
// The function takes individual of 40+1 dimensions and has two global optima in [1,1,...,1] and [0,0,...,0].
func ChuangF1(individual *base.BoolIndividual) []float64 {
	total, length := 0.0, individual.Len()
	chrom := individual.GetChromosome().([]bool)
	if chrom[length-1] == false {
		for i := 0; i < length-1; i += 4 {
			newInd := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			total += InvTrap(newInd)[0]
		}
	} else {
		for i := 0; i < length-1; i += 4 {
			newInd := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			total += Trap(newInd)[0]
		}
	}
	return []float64{total}
}

// ChuangF2 is maximization binary benchmark
//
// Binary deceptive function from : Multivariate Multi-Model Approach for Globally Multimodal Problems by Chung-Yao Chuang and Wen-Lian Hsu.
//
// The function takes individual of 40+1 dimensions and has four global optima in [1,1,...,0,0], [0,0,...,1,1], [1,1,...,1] and [0,0,...,0].
func ChuangF2(individual *base.BoolIndividual) []float64 {
	total, length := 0.0, individual.Len()
	chrom := individual.GetChromosome().([]bool)
	if chrom[length-2] == false && chrom[length-1] == false {
		for i := 0; i < length-2; i += 8 {
			newInd1 := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			newInd2 := base.NewBoolIndividual(chrom[i+4:i+8], individual.GetFitness().Clone())
			total += InvTrap(newInd1)[0] + InvTrap(newInd2)[0]
		}
	} else if chrom[length-2] == false && chrom[length-1] == true {
		for i := 0; i < length-2; i += 8 {
			newInd1 := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			newInd2 := base.NewBoolIndividual(chrom[i+4:i+8], individual.GetFitness().Clone())
			total += InvTrap(newInd1)[0] + Trap(newInd2)[0]
		}
	} else if chrom[length-2] == true && chrom[length-1] == false {
		for i := 0; i < length-2; i += 8 {
			newInd1 := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			newInd2 := base.NewBoolIndividual(chrom[i+4:i+8], individual.GetFitness().Clone())
			total += Trap(newInd1)[0] + InvTrap(newInd2)[0]
		}
	} else {
		for i := 0; i < length-2; i += 8 {
			newInd1 := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			newInd2 := base.NewBoolIndividual(chrom[i+4:i+8], individual.GetFitness().Clone())
			total += Trap(newInd1)[0] + Trap(newInd2)[0]
		}
	}
	return []float64{total}
}

// ChuangF3 is maximization binary benchmark
//
// Binary deceptive function from : Multivariate Multi-Model Approach for Globally Multimodal Problems by Chung-Yao Chuang and Wen-Lian Hsu.
//
// The function takes individual of 40+1 dimensions and has two global optima in [1,1,...,1] and [0,0,...,0].
func ChuangF3(individual *base.BoolIndividual) []float64 {
	total, length := 0.0, individual.Len()
	chrom := individual.GetChromosome().([]bool)
	if chrom[length-1] == false {
		for i := 0; i < length-1; i += 4 {
			newInd := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			total += InvTrap(newInd)[0]
		}
	} else {
		for i := 2; i < length-3; i += 4 {
			newInd1 := base.NewBoolIndividual(chrom[i:i+4], individual.GetFitness().Clone())
			total += InvTrap(newInd1)[0]
		}
		chrom1, chrom2 := chrom[length-2:], chrom[:2]
		newChrom := make([]bool, len(chrom1)+len(chrom2))
		copy(newChrom, chrom1)
		copy(newChrom[len(chrom1):], chrom2)
		total += Trap(base.NewBoolIndividual(newChrom, individual.GetFitness().Clone()))[0]
	}
	return []float64{total}
}

// RoyalRoad1 is binary benchmark
//
// Royal Road Function R1 as presented by Melanie Mitchell in : "An introduction to Genetic Algorithms".
func RoyalRoad1(individual *base.BoolIndividual, order int) []float64 {
	total, nelem := 0.0, individual.Len()
	maxValue := math.Pow(2.0, float64(order)) - 1
	for i := 0; i < nelem; i++ {
		value := 0.0
		chrom := individual.GetChromosome().([]bool)[i*order : i*order+order]
		for _, val := range chrom {
			if val {
				value *= 2.0
			}
		}
		total += float64(order) + value/maxValue
	}
	return []float64{total}
}

// RoyalRoad2 is binary benchmark
//
// Royal Road Function R2 as presented by Melanie Mitchell in : "An introduction to Genetic Algorithms".
func RoyalRoad2(individual *base.BoolIndividual, order int) []float64 {
	total := 0.0
	maxValue := int(math.Pow(2.0, float64(order)))
	for nOrder := order; nOrder < maxValue; nOrder *= 2 {
		total += RoyalRoad1(individual, nOrder)[0]
	}
	return []float64{total}
}
