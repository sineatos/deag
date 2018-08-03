package base

// Individual is a type which has the information of solution
type Individual interface {
	// Len returns the size of the chromosome
	Len() int
	// Clone is used to create a individual which has the same data
	Clone() interface{}
	// GetChromosome gets the chromosome slice (not a copy)
	GetChromosome() interface{}
	// SetChromosome sets the chromosome slice
	SetChromosome(chromosome interface{})
	// GetFitness returns the individual's fitness (not a copy)
	GetFitness() *Fitness
	// IsEqual returns if the other individual is equal to the individual
	IsEqual(other Individual) bool

	String() string
}

// Individuals is the slice of Individual
type Individuals []Individual

// ESIndividual is a Individual type which using in ES
type ESIndividual interface {
	Individual

	// SLen returns the size of strategies
	SLen() int
	// GetStrategies return strategies(not copy)
	GetStrategies() []float64
	// SetStrategies set strategies
	SetStrategies(strategies []float64)
}

// Len returns the size of Individuals
func (inds Individuals) Len() int {
	return len(inds)
}

// Less returns true if inds[i]'s fitness worse than inds[j]'s
func (inds Individuals) Less(i, j int) bool {
	return inds[i].GetFitness().Less(inds[j].GetFitness())
}

// Swap swaps two elements
func (inds Individuals) Swap(i, j int) {
	inds[i], inds[j] = inds[j], inds[i]
}
