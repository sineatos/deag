package benchmarks

import (
	"math"
	"testing"

	"github.com/sineatos/deag/base"
)

const (
	dims int     = 20
	min  float64 = -1.0
	max  float64 = 1.0
	eps  float64 = 1E-10
)

// fastGenerateFloat64Individual returns a individual initial with specifed
func fastGenerateFloat64Individual(dim int, v, w float64) *base.Float64Individual {
	values := make([]float64, dim)
	weights := []float64{w}
	for i := range values {
		values[i] = v
	}
	weights[0] = w
	fitness := base.NewFitness(weights)
	individual := base.NewFloat64Individual(values, fitness)
	return individual
}

// checkValue checks the target if equals to actual value
func checkValue(t *testing.T, target, value float64, dim int) {
	if math.Abs(value-target) >= eps {
		t.Errorf("Expected %v, got %v in dim=%v", target, value, dim)
	}
}

func TestPlane(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Plane(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestSphere(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Sphere(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestCigar(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Cigar(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestRosenbrock(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 1, min)
		target, value := 0.0, Rosenbrock(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestH1(t *testing.T) {
	individual := fastGenerateFloat64Individual(2, 1, max)
	chrom := individual.GetChromosome().([]float64)
	chrom[0], chrom[1] = 8.6998, 6.7665
	individual.SetChromosome(chrom)
	target, value := 2.0, H1(individual)[0]
	checkValue(t, target, value, 2)
}

func TestAckley(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Ackley(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestBohachevsky(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Bohachevsky(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestRastrigin(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Rastrigin(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestRastriginScaled(t *testing.T) {
	for dim := 2; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, RastriginScaled(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestRastriginSkew(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, RastriginSkew(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestSchaffer(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 0, min)
		target, value := 0.0, Schaffer(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestSchwefel(t *testing.T) {
	for dim := 1; dim < dims; dim++ {
		individual := fastGenerateFloat64Individual(dim, 420.96874636, min)
		target, value := 0.0, Schwefel(individual)[0]
		checkValue(t, target, value, dim)
	}
}

func TestHimmelblau(t *testing.T) {
	individual := fastGenerateFloat64Individual(2, 1, min)
	chrom := individual.GetChromosome().([]float64)
	chrom[0], chrom[1] = 3.0, 2.0
	individual.SetChromosome(chrom)
	target, value := 0.0, Himmelblau(individual)[0]
	if math.Abs(value-target) >= eps {
		t.Errorf("Expected %v, got %v of %v", target, value, chrom)
	}

	chrom[0], chrom[1] = -2.805118, 3.131312
	individual.SetChromosome(chrom)
	target, value = 0.0, Himmelblau(individual)[0]
	if math.Abs(value-target) >= eps {
		t.Errorf("Expected %v, got %v of %v", target, value, chrom)
	}

	chrom[0], chrom[1] = -3.779310, -3.283186
	individual.SetChromosome(chrom)
	target, value = 0.0, Himmelblau(individual)[0]
	if math.Abs(value-target) >= eps {
		t.Errorf("Expected %v, got %v of %v", target, value, chrom)
	}

	chrom[0], chrom[1] = 3.584428, -1.848126
	individual.SetChromosome(chrom)
	target, value = 0.0, Himmelblau(individual)[0]
	if math.Abs(value-target) >= eps {
		t.Errorf("Expected %v, got %v of %v", target, value, chrom)
	}
}

func TestShekel(t *testing.T) {
	// The Shekel function has not tested
}
