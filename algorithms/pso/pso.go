package pso

import (
	"math/rand"

	"github.com/sineatos/deag/base"
	"github.com/sineatos/deag/benchmarks"
	"github.com/sineatos/deag/tools/support"
)

// PSO is the standard particle swarm optimaization
type PSO struct {
	population base.Individuals
	c1         float64
	c2         float64
	gbest      *Particle
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

// NewPSO returns *PSO.
// c1 is coefficient of local information.
// c2 is coefficient of global information
// maxGen is maximum running generation.
// maxFES is maximum function evalutions, maxFES < 0 means maxFES=INF.
// stat is Statistics, optional.
// hof is HallOfFame records the best individuals, optional.
// evaluator is a Float64Evaluator.
func NewPSO(c1, c2 float64, maxGen, maxFES int, stat support.Statistics, hof support.HallOfFame, evaluator benchmarks.Float64Evaluator) *PSO {
	return &PSO{
		c1:        c1,
		c2:        c2,
		maxGen:    maxGen,
		maxFES:    maxFES,
		stat:      stat,
		hof:       hof,
		evaluator: evaluator,
	}
}

// Init initializes the population and prepared for some data
func (evol *PSO) Init(population base.Individuals) {
	evol.currentFES = 0
	evol.gen = 0
	evol.logbook = support.NewDefaultLogbook(evol.maxGen, 0)
	evol.population = population
	evol.size = population.Len()

	if !evol.IsTerminated() {
		evol.evaluate(evol.population)
		evol.log()
	}
}

// IsTerminated returns if the evolution is terminated
func (evol *PSO) IsTerminated() (flag bool) {
	flag = evol.gen >= evol.maxGen
	flag = flag || (evol.maxFES > 0 && evol.currentFES+evol.size > evol.maxFES)
	flag = flag || (evol.hof.Len() > 0 && evol.hof.Get(0).GetFitness().GetValues()[0] < 1E-14)
	return
}

// Evolve runs the evol a time/generation per call and return generation time
func (evol *PSO) Evolve() interface{} {
	evol.gen++
	if !evol.IsTerminated() {
		// create offsprings (only clone)
		offsprings := make(base.Individuals, evol.size)
		gBest := evol.hof.Get(0)
		tGBest := gBest
		for i, ind := range evol.population {
			cInd := ind.Clone().(*Particle)
			speed := cInd.GetSpeed()
			chrom := cInd.GetChromosome().([]float64)
			pChrom := cInd.GetPBest().GetChromosome().([]float64)
			gChrom := gBest.GetChromosome().([]float64)
			for j := range speed {
				speed[j] += evol.c1*rand.Float64()*(pChrom[j]-chrom[j]) + evol.c2*rand.Float64()*(gChrom[j]-chrom[j])
				chrom[j] += speed[j]
			}
			cInd.SetSpeed(speed)
			cInd.GetFitness().SetValues(evol.evaluator(&cInd.Float64Individual))
			if cInd.GetFitness().Greater(ind.GetFitness()) {
				if cInd.GetFitness().Greater(cInd.GetPBest().GetFitness()) {
					cInd.SetPBest(cInd)
					if cInd.GetFitness().Greater(tGBest.GetFitness()) {
						tGBest = cInd
					}
				}
				offsprings[i] = cInd
			} else {
				offsprings[i] = ind.Clone().(*Particle)
			}
		}
		gBest = tGBest
		evol.population = offsprings
		evol.currentFES += evol.size
		// log
		evol.log()
	}
	return evol.gen
}

// Run executes Evolve() until the terminal condition satisfied
func (evol *PSO) Run() {
	for !evol.IsTerminated() {
		evol.Evolve()
	}
}

// GetLogbook returns the logbook saving data
func (evol *PSO) GetLogbook() support.Logbook {
	return evol.logbook
}

// GetHallOfFame returns the HallOfFame saving best individuals
func (evol *PSO) GetHallOfFame() support.HallOfFame {
	return evol.hof
}

func (evol *PSO) evaluate(Individuals base.Individuals) {
	for _, ind := range evol.population {
		fInd := ind.(*Particle)
		fit := evol.evaluator(&fInd.Float64Individual)
		fInd.GetFitness().SetValues(fit)
		fInd.SetPBest(fInd)
	}
	evol.currentFES += evol.size
}

func (evol *PSO) log() {
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
