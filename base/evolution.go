package base

// Evolution provides the action which evolution algorithm or other mega-heuristic algorithm can execute
type Evolution interface {
	// Init initializes using population and prepared for some data
	Init(population Individuals)
	// IsTerminated returns if the evolution is terminated
	IsTerminated() bool
	// Evolve runs the evolution a time/generation per call and return some information about evolution
	Evolve() interface{}
}
