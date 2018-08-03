package constraint

import (
	"github.com/sineatos/deag/base"
	"testing"
)

func TestClosestValidPenalty(t *testing.T) {
	minBound := make([]float64, 30)
	maxBound := make([]float64, 30)
	for i := range maxBound {
		maxBound[i] = 1.0
	}

	distance := func(fInd, oInd *base.Float64Individual) float64 {
		fChrom := fInd.GetChromosome().([]float64)
		oChrom := oInd.GetChromosome().([]float64)
		ans := 0.0
		for i, f := range fChrom {
			ans += (f - oChrom[i]) * (f - oChrom[i])
		}
		return ans
	}

	closestFeasible := func(individual *base.Float64Individual) *base.Float64Individual {
		fInd := individual.Clone().(*base.Float64Individual)
		fChrom := fInd.GetChromosome().([]float64)
		for i, f := range fChrom {
			if f < minBound[i] {
				f = minBound[i]
			}
			if f > maxBound[i] {
				f = maxBound[i]
			}
			fChrom[i] = f
		}
		return fInd
	}

	valid := func(individual *base.Float64Individual) bool {
		chrom := individual.GetChromosome().([]float64)
		for i, c := range chrom {
			if c < minBound[i] || c > maxBound[i] {
				return false
			}
		}
		return true
	}

	zdt2 := func(individual *base.Float64Individual) []float64 {
		chrom := individual.GetChromosome().([]float64)
		sum := 0.0
		for _, c := range chrom[1:] {
			sum += c
		}
		g := 1.0 + 9.0*sum/float64(len(chrom)-1)
		f1 := chrom[0]
		f2 := g * (1.0 - (f1/g)*(f1/g))
		return []float64{f1, f2}
	}
	con := NewClosestValidPenalty(valid, closestFeasible, 1.0e-6, distance, zdt2)

	ind1 := base.NewFloat64Individual([]float64{-5.6468535666e-01, 2.2483050478e+00, -1.1087909644e+00, -1.2710112861e-01, 1.1682438733e+00, -1.3642007438e+00, -2.1916417835e-01, -5.9137308999e-01, -1.0870160336e+00, 6.0515070232e-01, 2.1532075914e+00, -2.6164718271e-01, 1.5244071578e+00, -1.0324305612e+00, 1.2858152343e+00, -1.2584683962e+00, 1.2054392372e+00, -1.7429571973e+00, -1.3517256013e-01, -2.6493429355e+00, -1.3051320798e-01, 2.2641961090e+00, -2.5027232340e+00, -1.2844874148e+00, 1.9955852925e+00, -1.2942218834e+00, 3.1340109155e+00, 1.6440111097e+00, -1.7750105857e+00, 7.7610242710e-01}, base.NewFitness([]float64{-1.0, -1.0}))

	// 4.142787904187075e-05 4.532154468043869
	t.Log(con.AdjustAndEvolve(ind1))
	// false
	t.Log("Individuals is valid: ", valid(ind1))
}
