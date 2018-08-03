package de

import (
	"math/rand"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
	"github.com/sineatos/deag/tools/support"
)

// Rand1 is implement of using DE/rand/1 as mutation operator
type Rand1 struct {
	population base.Individuals
	f          float64
	cr         float64
	stat       support.Statistics
	hof        support.HallOfFame
	logbook    support.Logbook
	size       int
	maxFES     int
	currentFES int
	gen        int
	maxGen     int
	evaluator  benchmarks.Float64Evaluator
}

// NewDERand1 returns *Rand1.
// f is scale factor.
// cr is crossover rate.
// maxGen is maximum running generation.
// maxFES is maximum function evalutions, maxFES < 0 means maxFES=INF.
// stat is Statistics, optional.
// hof is HallOfFame records the best individuals, optional.
// evaluator is a Float64Evaluator.
func NewDERand1(f, cr float64, maxGen, maxFES int, stat support.Statistics, hof support.HallOfFame, evaluator benchmarks.Float64Evaluator) *Rand1 {
	return &Rand1{
		f:         f,
		cr:        cr,
		maxGen:    maxGen,
		maxFES:    maxFES,
		stat:      stat,
		hof:       hof,
		evaluator: evaluator,
	}
}

// Init initializes the population and prepared for some data
func (evol *Rand1) Init(population base.Individuals) {
	evol.currentFES = 0
	evol.gen = 0
	evol.logbook = support.NewDefaultLogbook(evol.maxGen, 0)
	evol.population = population
	evol.size = population.Len()

	if evol.hof == nil {
		evol.hof = support.NewDefaultHallOfFame(1, nil)
	}

	if !evol.IsTerminated() {
		evol.evaluate(evol.population)
		evol.log()
	}
}

// IsTerminated returns if the evolution is terminated
func (evol *Rand1) IsTerminated() (flag bool) {
	flag = evol.gen >= evol.maxGen
	flag = flag || (evol.maxFES > 0 && evol.currentFES+evol.size > evol.maxFES)
	flag = flag || (evol.hof.Len() > 0 && evol.hof.Get(0).GetFitness().GetValues()[0] < 1E-14)
	return
}

// Evolve runs the evol a time/generation per call and return generation time
func (evol *Rand1) Evolve() interface{} {
	evol.gen++
	if !evol.IsTerminated() {
		// create offsprings (only clone)
		offsprings := make(base.Individuals, evol.size)
		for i, agent := range evol.population {
			// prepare data
			// select two different individuals and they are different from agent
			inds := rand.Perm(evol.population.Len())
			rAmounts := 3
			rChroms, rLoc := make([][]float64, rAmounts), 0
			for _, x := range inds {
				if agent != evol.population[x] {
					rChroms[rLoc] = evol.population[x].GetChromosome().([]float64)
					rLoc++
					if rLoc == rAmounts {
						break
					}
				}
			}

			t := agent.Clone().(*base.Float64Individual)
			tChrom := t.GetChromosome().([]float64)
			agentChrom := agent.GetChromosome().([]float64)
			// mutation
			for j, ch := range rChroms[0] {
				tChrom[j] = ch + evol.f*(rChroms[1][j]-rChroms[2][j])
			}
			// crossover
			index := rand.Intn(agent.Len())
			for j, value := range agentChrom {
				if index == j || rand.Float64() < evol.cr {
					tChrom[j] = value
				}
			}
			t.GetFitness().SetValues(evol.evaluator(t))
			// selection
			if t.GetFitness().Greater(agent.GetFitness()) {
				offsprings[i] = t
			} else {
				offsprings[i] = evol.population[i].Clone().(base.Individual)
			}
		}
		evol.population = offsprings
		evol.currentFES += evol.size
		// log
		evol.log()
	}
	return evol.gen
}

// Run executes Evolve() until the terminal condition satisfied
func (evol *Rand1) Run() {
	for !evol.IsTerminated() {
		evol.Evolve()
	}
}

// GetLogbook returns the logbook saving data
func (evol *Rand1) GetLogbook() support.Logbook {
	return evol.logbook
}

// GetHallOfFame returns the HallOfFame saving best individuals
func (evol *Rand1) GetHallOfFame() support.HallOfFame {
	return evol.hof
}

func (evol *Rand1) evaluate(Individuals base.Individuals) {
	for _, ind := range evol.population {
		fInd := ind.(*base.Float64Individual)
		fit := evol.evaluator(fInd)
		fInd.GetFitness().SetValues(fit)
	}
	evol.currentFES += evol.size
}

func (evol *Rand1) log() {
	var datas support.Dict
	if evol.stat != nil {
		datas = evol.stat.Compile(evol.population)
	} else {
		datas = make(support.Dict, 3)
	}
	datas[support.GEN] = evol.gen
	datas[support.FES] = evol.size
	evol.logbook.Record(datas)

	evol.hof.Update(evol.population)
}
