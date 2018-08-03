package support

import (
	"math"

	"github.com/sineatos/deag/base"
)

const (
	// GEN is the name of statistics which is the generation number
	GEN string = "gen"
	// FES is the name of statistics which is the function evalutions
	FES string = "FES"
)

// StatFunction is type of function registered in Statistics
type StatFunction func(interface{}) interface{}

// Statistics is an interface of statistics
type Statistics interface {
	// GetName returns the name of Statistics
	GetName() string
	// GetConverter returns the result which is converted from input datas
	GetConverter() StatFunction
	// Register registers the function with name
	Register(name string, function StatFunction)
	// Compile calculates Statistics according to input
	Compile(input interface{}) Dict
}

// MultiStatistics is an interface of multistatistics which is the integrate of Statistics
type MultiStatistics interface {
	Statistics

	// AddStats adds stat and returns the name of stat
	AddStats(stat Statistics) string
	// RemoveStats remove Statistics according to name
	RemoveStats(name string)
	// GetStats returns Statistics according to name
	GetStats(name string) Statistics
	// GetAllStats returns all Statistics
	GetAllStats() map[string]Statistics
}

// DefaultStatistics is a type of Statistics which calculating the statistics of single objective fitness
type DefaultStatistics struct {
	name      string
	convertor StatFunction
	functions map[string]StatFunction
}

// NewDefaultStatistics returns a default Statistics with name
func NewDefaultStatistics(name string, convertor StatFunction) *DefaultStatistics {
	functions := make(map[string]StatFunction)
	if convertor == nil {
		convertor = func(input interface{}) interface{} { return input }
	}
	return &DefaultStatistics{name: name, functions: functions, convertor: convertor}
}

// GetName returns the name of Statistics
func (stat *DefaultStatistics) GetName() string {
	return stat.name
}

// GetConverter returns the result which is converted from input datas
func (stat *DefaultStatistics) GetConverter() StatFunction {
	return stat.convertor
}

// Register registers function with name
func (stat *DefaultStatistics) Register(name string, function StatFunction) {
	stat.functions[name] = function
}

// Compile calculates Statistics according to input
func (stat *DefaultStatistics) Compile(input interface{}) Dict {
	datas := stat.convertor(input)
	ans := make(Dict, len(stat.functions))
	for key, function := range stat.functions {
		ans[key] = function(datas)
	}
	return ans
}

// DefaultMultiStatistics is a type of Statistics which calculating the statistics of single objective fitness
type DefaultMultiStatistics struct {
	DefaultStatistics
	dict map[string]Statistics
}

// NewDefaultMultiStatistics returns a default MultiStatistics with name
func NewDefaultMultiStatistics(name string) *DefaultMultiStatistics {
	dict := make(map[string]Statistics)
	return &DefaultMultiStatistics{DefaultStatistics: *NewDefaultStatistics(name, nil), dict: dict}
}

// AddStats adds stat and returns the name of stat
func (mStat *DefaultMultiStatistics) AddStats(stat Statistics) string {
	mStat.dict[stat.GetName()] = stat
	return stat.GetName()
}

// RemoveStats remove Statistics according to name
func (mStat *DefaultMultiStatistics) RemoveStats(name string) {
	delete(mStat.dict, name)
}

// GetStats returns Statistics according to name
func (mStat *DefaultMultiStatistics) GetStats(name string) Statistics {
	return mStat.dict[name]
}

// GetAllStats returns all Statistics
func (mStat *DefaultMultiStatistics) GetAllStats() map[string]Statistics {
	ans := make(map[string]Statistics, len(mStat.dict))
	for k, v := range mStat.dict {
		ans[k] = v
	}
	return ans
}

// Register registers function with name
func (mStat *DefaultMultiStatistics) Register(name string, function StatFunction) {
	for _, stat := range mStat.dict {
		stat.Register(name, function)
	}
}

// Compile calculates Statistics according to input
func (mStat *DefaultMultiStatistics) Compile(input interface{}) Dict {
	ans := make(Dict, len(mStat.dict))
	for key, stat := range mStat.dict {
		ans[key] = stat.Compile(input)
	}
	return ans
}

// NewStatisticsBasedOnFitness returns a Statistics based on fitness with name, if name is "", the name default is "fitness"
func NewStatisticsBasedOnFitness(name string) Statistics {
	if name == "" {
		name = "fitness"
	}
	return NewDefaultStatistics(name, func(input interface{}) interface{} {
		inds := input.(base.Individuals)
		fitness := make([]float64, inds.Len())
		for i, ind := range inds {
			fitness[i] = ind.GetFitness().GetValues()[0]
		}
		return fitness
	})
}

// NewStatisticsBasedOnMOFitness returns a Statistics based on multi objective fitness with name, if name is "", the name default is "mo-fitness"
func NewStatisticsBasedOnMOFitness(name string) Statistics {
	if name == "" {
		name = "mo-fitness"
	}
	return NewDefaultStatistics(name, func(input interface{}) interface{} {
		inds := input.(base.Individuals)
		fitness := make([][]float64, inds.Len())
		for i, ind := range inds {
			fitness[i] = ind.GetFitness().GetValues()
		}
		return fitness
	})
}

// NewStatisticsBasedOnIndividual returns a Statistics based on individual with name, if name is "", the name default is "individual"
func NewStatisticsBasedOnIndividual(name string) Statistics {
	if name == "" {
		name = "individual"
	}
	return NewDefaultStatistics(name, func(input interface{}) interface{} {
		return input
	})
}

// StatFitnessMin is a StatFunction which input is []float64 and returns the slice's minimum
func StatFitnessMin(input interface{}) interface{} {
	fitness := input.([]float64)
	minF := math.MaxFloat64
	for _, fit := range fitness {
		minF = math.Min(minF, fit)
	}
	return minF
}

// StatFitnessMax is a StatFunction which input is []float64 and returns the slice's maximum
func StatFitnessMax(input interface{}) interface{} {
	fitness := input.([]float64)
	maxF := -math.MaxFloat64
	for _, fit := range fitness {
		maxF = math.Max(maxF, fit)
	}
	return maxF
}

// StatFitnessAvg is a StatFunction which input is []float64 and returns the slice's average
func StatFitnessAvg(input interface{}) interface{} {
	fitness := input.([]float64)
	avg := 0.0
	for _, fit := range fitness {
		avg += fit
	}
	return avg / float64(len(fitness))
}

// StatFitnessStd is a StatFunction which input is []float64 and returns the slice's standard deviation
func StatFitnessStd(input interface{}) interface{} {
	fitness := input.([]float64)
	avg, sum := 0.0, 0.0
	for _, fit := range fitness {
		avg += fit
	}
	avg /= float64(len(fitness))
	for _, fit := range fitness {
		sum += (fit - avg) * (fit - avg)
	}
	return math.Sqrt(sum / float64(len(fitness)))
}

// StatMOFitnessMin is a StatFunction which input is [][]float64 and returns the slice's minimum([]float64)
func StatMOFitnessMin(input interface{}) interface{} {
	fitness := input.([][]float64)
	minF := make([]float64, len(fitness[0]))
	for i := range minF {
		minF[i] = math.MaxFloat64
	}
	for _, fit := range fitness {
		for i, m := range minF {
			minF[i] = math.Min(m, fit[i])
		}
	}
	return minF
}

// StatMOFitnessMax is a StatFunction which input is [][]float64 and returns the slice's maximum([]float64)
func StatMOFitnessMax(input interface{}) interface{} {
	fitness := input.([][]float64)
	maxF := make([]float64, len(fitness[0]))
	for i := range maxF {
		maxF[i] = -math.MaxFloat64
	}
	for _, fit := range fitness {
		for i, m := range maxF {
			maxF[i] = math.Max(m, fit[i])
		}
	}
	return maxF
}

// StatMOFitnessAvg is a StatFunction which input is [][]float64 and returns the slice's average([]float64)
func StatMOFitnessAvg(input interface{}) interface{} {
	fitness := input.([][]float64)
	avg := make([]float64, len(fitness[0]))
	for _, fit := range fitness {
		for i := range avg {
			avg[i] += fit[i]
		}
	}
	size := float64(len(fitness))
	for i := range avg {
		avg[i] /= size
	}
	return avg
}

// StatMOFitnessStd is a StatFunction which input is [][]float64 and returns the slice's standard deviation([]float64)
func StatMOFitnessStd(input interface{}) interface{} {
	fitness := input.([][]float64)
	avg := make([]float64, len(fitness[0]))
	sum := make([]float64, len(fitness[0]))
	for _, fit := range fitness {
		for i := range avg {
			avg[i] += fit[i]
		}
	}
	size := float64(len(fitness))
	for i := range avg {
		avg[i] /= size
	}
	for _, fit := range fitness {
		for i, a := range avg {
			sum[i] += math.Pow(fit[i]-a, 2.0)
		}
	}
	for i, s := range sum {
		sum[i] = math.Sqrt(s / size)
	}
	return sum
}
