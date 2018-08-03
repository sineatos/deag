package benchmarks

import (
	"math"

	"github.com/sineatos/deag/base"
)

// Kursawe multiobjective function.
//
// `f_{\\text{Kursawe}1}(\\mathbf{x}) = \\sum_{i=1}^{N-1} -10 e^{-0.2 \\sqrt{x_i^2 + x_{i+1}^2} }`
//
// `f_{\\text{Kursawe}1}(\\mathbf{x}) = \\sum_{i=1}^{N-1} -10 e^{-0.2 \\sqrt{x_i^2 + x_{i+1}^2} }`
func Kursawe(individual *base.Float64Individual) []float64 {
	f1Ans, f2Ans := 0.0, 0.0
	chrom := individual.GetChromosome().([]float64)
	f2Ans = math.Pow(math.Abs(chrom[0]), 0.8) + 5*math.Sin(math.Pow(chrom[0], 3.0))
	for i, ch := range chrom[1:] {
		f1Ans += -10.0 * math.Exp(-0.2*float64(chrom[i]*chrom[i]+ch*ch))
		f2Ans += math.Pow(math.Abs(chrom[i]), 0.8) + 5*math.Sin(math.Pow(chrom[i], 3.0))
	}
	return []float64{f1Ans, f2Ans}
}

// SchafferMo is Schaffer's multiobjective function on a one attribute individual.
// From: J. D. Schaffer, "Multiple objective optimization with vector From: J. D. Schaffer, "Multiple objective optimization with vector Conference on Genetic Algorithms, 1987.
//
// `f_{\\text{Schaffer}1}(\\mathbf{x}) = x_1^2`
//
// `f_{\\text{Schaffer}1}(\\mathbf{x}) = x_1^2`
func SchafferMo(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	return []float64{math.Pow(chrom[0], 2.0), math.Pow(chrom[0]-2.0, 2.0)}
}

// ZDT1 multiobjective function.
//
// `g(\\mathbf{x}) = 1 + \\frac{9}{n-1}\\sum_{i=2}^n x_i`
//
// `f_{\\text{ZDT1}1}(\\mathbf{x}) = x_1`
//
// `f_{\\text{ZDT1}2}(\\mathbf{x}) = g(\\mathbf{x})\\left[1 - \\sqrt{\\frac{x_1}{g(\\mathbf{x})}}\\right]`
func ZDT1(individual *base.Float64Individual) []float64 {
	g := 0.0
	chrom := individual.GetChromosome().([]float64)
	for _, ch := range chrom[1:] {
		g += ch
	}
	g = 1.0 + 9.0*g/float64(len(chrom)-1)
	f1 := chrom[0]
	f2 := g * (1.0 - math.Sqrt(f1/g))
	return []float64{f1, f2}
}

// ZDT2 multiobjective function.
//
// `g(\\mathbf{x}) = 1 + \\frac{9}{n-1}\\sum_{i=2}^n x_i`
//
// `f_{\\text{ZDT2}1}(\\mathbf{x}) = x_1`
//
// `f_{\\text{ZDT2}2}(\\mathbf{x}) = g(\\mathbf{x})\\left[1 - \\left(\\frac{x_1}{g(\\mathbf{x})}\\right)^2\\right]`
func ZDT2(individual *base.Float64Individual) []float64 {
	g := 0.0
	chrom := individual.GetChromosome().([]float64)
	for _, ch := range chrom[1:] {
		g += ch
	}
	g = 1.0 + 9.0*g/float64(len(chrom)-1)
	f1 := chrom[0]
	f2 := g * (1.0 - math.Pow(f1/g, 2.0))
	return []float64{f1, f2}
}

// ZDT3 multiobjective function.
//
// `g(\\mathbf{x}) = 1 + \\frac{9}{n-1}\\sum_{i=2}^n x_i`
//
// `f_{\\text{ZDT3}1}(\\mathbf{x}) = x_1`
//
// `f_{\\text{ZDT3}2}(\\mathbf{x}) = g(\\mathbf{x})\\left[1 - \\sqrt{\\frac{x_1}{g(\\mathbf{x})}} - \\frac{x_1}{g(\\mathbf{x})}\\sin(10\\pi x_1)\\right]`
func ZDT3(individual *base.Float64Individual) []float64 {
	g := 0.0
	chrom := individual.GetChromosome().([]float64)
	for _, ch := range chrom[1:] {
		g += ch
	}
	g = 1.0 + 9.0*g/float64(len(chrom)-1)
	f1 := chrom[0]
	f2 := g * (1.0 - math.Sqrt(f1/g) - f1/g*math.Sin(10.0*math.Pi*f1))
	return []float64{f1, f2}
}

// ZDT4 multiobjective function.
//
// `g(\\mathbf{x}) = 1 + 10(n-1) + \\sum_{i=2}^n \\left[ x_i^2 - 10\\cos(4\\pi x_i) \\right]`
//
// `f_{\\text{ZDT4}1}(\\mathbf{x}) = x_1`
//
// `f_{\\text{ZDT4}2}(\\mathbf{x}) = g(\\mathbf{x})\\left[ 1 - \\sqrt{x_1/g(\\mathbf{x})} \\right]`
func ZDT4(individual *base.Float64Individual) []float64 {
	g := 0.0
	chrom := individual.GetChromosome().([]float64)
	for _, ch := range chrom[1:] {
		g += ch*ch - 10.0*math.Cos(4.0*math.Pi*ch)
	}
	g += 1.0 + 10.0*float64(len(chrom)-1)
	f1 := chrom[0]
	f2 := g * (1.0 - math.Sqrt(f1/g))
	return []float64{f1, f2}
}

// ZDT6 multiobjective function.
//
// `g(\\mathbf{x}) = 1 + 9 \\left[ \\left(\\sum_{i=2}^n x_i\\right)/(n-1) \\right]^{0.25}`
//
// `f_{\\text{ZDT6}1}(\\mathbf{x}) = 1 - \\exp(-4x_1)\\sin^6(6\\pi x_1)`
//
// `f_{\\text{ZDT6}2}(\\mathbf{x}) = g(\\mathbf{x}) \left[ 1 - (f_{\\text{ZDT6}1}(\\mathbf{x})/g(\\mathbf{x}))^2 \\right]`
func ZDT6(individual *base.Float64Individual) []float64 {
	g := 0.0
	chrom := individual.GetChromosome().([]float64)
	for _, ch := range chrom[1:] {
		g += ch
	}
	g += 1.0 + 9.0*math.Pow(g/float64(len(chrom)-1), 0.25)
	f1 := 1.0 - math.Exp(-4.0*chrom[0])*math.Pow(6.0*math.Pi*chrom[0], 6.0)
	f2 := g * (1.0 - math.Pow(f1/g, 2.0))
	return []float64{f1, f2}
}

// DTLZ1 multiobjective function. It returns a slice of obj values.
//
// The individual must have at least obj elements. From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
//
// `g(\\mathbf{x}_m) = 100\\left(|\\mathbf{x}_m| + \sum_{x_i \in \\mathbf{x}_m}\\left((x_i - 0.5)^2 - \\cos(20\pi(x_i - 0.5))\\right)\\right)`
//
// `f_{\\text{DTLZ1}1}(\\mathbf{x}) = \\frac{1}{2} (1 + g(\\mathbf{x}_m)) \\prod_{i=1}^{m-1}x_i`
//
// `f_{\\text{DTLZ1}2}(\\mathbf{x}) = \\frac{1}{2} (1 + g(\\mathbf{x}_m)) (1-x_{m-1}) \\prod_{i=1}^{m-2}x_i`
//
// `\\ldots`
//
// `f_{\\text{DTLZ1}m-1}(\\mathbf{x}) = \\frac{1}{2} (1 + g(\\mathbf{x}_m)) (1 - x_2) x_1`
//
// `f_{\\text{DTLZ1}m}(\\mathbf{x}) = \\frac{1}{2} (1 - x_1)(1 + g(\\mathbf{x}_m))`
//
// Where `m` is the number of objectives and `\\mathbf{x}_m` is a vector of the remaining attributes `[x_m~\\ldots~x_n]` of the individual in `n > m` dimensions.
func DTLZ1(individual *base.Float64Individual, obj int) []float64 {
	chrom := individual.GetChromosome().([]float64)
	g := float64(len(chrom[obj-1:]))
	for _, ch := range chrom[obj-1:] {
		g += math.Pow(ch-0.5, 2.0) - math.Cos(20.0*math.Pi*(ch-0.5))
	}
	g *= 100.0
	f, pr := make([]float64, obj), 1.0
	for _, ch := range chrom[:obj-1] {
		f[0] += pr * ch
		pr = f[0]
	}
	f[0] *= 0.5 * (1.0 + g)
	j := 1
	for m := obj - 2; m >= 0; m-- {
		pr = 1.0
		for _, ch := range chrom[:m] {
			f[j] += pr * ch
			pr = f[0]
		}
		f[j] *= 0.5
		f[j] *= (1.0 - chrom[m]) * (1.0 + g)
		j++
	}
	return f
}

// DTLZ2 multiobjective function. It returns a slice of obj values.
//
// The individual must have at least obj elements.
//
// From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
//
// `g(\\mathbf{x}_m) = \\sum_{x_i \in \\mathbf{x}_m} (x_i - 0.5)^2`
//
// `f_{\\text{DTLZ2}1}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\prod_{i=1}^{m-1} \\cos(0.5x_i\pi)`
//
// `f_{\\text{DTLZ2}2}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\sin(0.5x_{m-1}\pi ) \\prod_{i=1}^{m-2} \\cos(0.5x_i\pi)`
//
// `\\ldots`
//
// `f_{\\text{DTLZ2}m}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\sin(0.5x_{1}\pi )`
//
// Where `m` is the number of objectives and `\\mathbf{x}_m` is a vector of the remaining attributes `[x_m~\\ldots~x_n]` of the individual in `n > m` dimensions.
func DTLZ2(individual *base.Float64Individual, obj int) []float64 {
	chrom := individual.GetChromosome().([]float64)
	xc, xm := chrom[:obj-1], chrom[obj-1:]
	g := 0.0
	for _, ch := range xm {
		g += math.Pow(ch-0.5, 2.0)
	}
	f, pr := make([]float64, obj), 1.0
	for _, ch := range xc {
		f[0] += pr * math.Cos(math.Pi*ch*0.5)
		pr = f[0]
	}
	f[0] *= (1.0 + g)
	j := 1
	for m := obj - 2; m >= 0; m-- {
		pr = 1.0
		for _, ch := range xc[:m] {
			f[j] += pr * math.Cos(math.Pi*ch*0.5)
			pr = f[j]
		}
		f[j] *= (1.0 - chrom[m]) * (1.0 + g) * math.Sin(0.5*xc[m]*math.Pi)
		j++
	}
	return f
}

// DTLZ3 multiobjective function. It returns a slice of obj values.
//
// The individual must have at least obj elements.
//
// From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
//
// `g(\\mathbf{x}_m) = 100\\left(|\\mathbf{x}_m| + \sum_{x_i \in \\mathbf{x}_m}\\left((x_i - 0.5)^2 - \\cos(20\pi(x_i - 0.5))\\right)\\right)`
//
// `f_{\\text{DTLZ3}1}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\prod_{i=1}^{m-1} \\cos(0.5x_i\pi)`
//
// `f_{\\text{DTLZ3}2}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\sin(0.5x_{m-1}\pi ) \\prod_{i=1}^{m-2} \\cos(0.5x_i\pi)`
//
// `\\ldots`
//
// `f_{\\text{DTLZ3}m}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\sin(0.5x_{1}\pi )`
//
// Where `m` is the number of objectives and `\\mathbf{x}_m` is a vector of the remaining attributes `[x_m~\\ldots~x_n]` of the individual in `n > m` dimensions.
func DTLZ3(individual *base.Float64Individual, obj int) []float64 {
	chrom := individual.GetChromosome().([]float64)
	xc, xm := chrom[:obj-1], chrom[obj-1:]
	g := float64(len(xm))
	for _, ch := range xm {
		g += math.Pow(ch-0.5, 2.0) - math.Cos(20.0*math.Pi*(ch-0.5))
	}
	g *= 100.0
	f, pr := make([]float64, obj), 1.0
	for _, ch := range xc {
		f[0] += pr * math.Cos(math.Pi*ch*0.5)
		pr = f[0]
	}
	f[0] *= (1.0 + g)
	j := 1
	for m := obj - 2; m >= 0; m-- {
		pr = 1.0
		for _, ch := range xc[:m] {
			f[j] += pr * math.Cos(math.Pi*ch*0.5)
			pr = f[j]
		}
		f[j] *= (1.0 - chrom[m]) * (1.0 + g) * math.Sin(0.5*xc[m]*math.Pi)
		j++
	}
	return f
}

// DTLZ4 multiobjective function. It returns a slice of obj values.
//
// The individual must have at least obj elements. The alpha parameter allows for a meta-variable mapping in `dtlz2` `x_i \\rightarrow x_i^\\alpha`, the authors suggest `\\alpha = 100`.
//
// From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
//
// `g(\\mathbf{x}_m) = \\sum_{x_i \in \\mathbf{x}_m} (x_i - 0.5)^2`
//
// `f_{\\text{DTLZ4}1}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\prod_{i=1}^{m-1} \\cos(0.5x_i^\\alpha\pi)`
//
// `f_{\\text{DTLZ4}2}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\sin(0.5x_{m-1}^\\alpha\pi ) \\prod_{i=1}^{m-2} \\cos(0.5x_i^\\alpha\pi)`
//
// `\\ldots`
//
// `f_{\\text{DTLZ4}m}(\\mathbf{x}) = (1 + g(\\mathbf{x}_m)) \\sin(0.5x_{1}^\\alpha\pi )`
//
// Where `m` is the number of objectives and `\\mathbf{x}_m` is a vector of the remaining attributes `[x_m~\\ldots~x_n]` of the individual in `n > m` dimensions.
func DTLZ4(individual *base.Float64Individual, obj int, alpha float64) []float64 {
	chrom := individual.GetChromosome().([]float64)
	xc, xm := chrom[:obj-1], chrom[obj-1:]
	g := 0.0
	for _, ch := range xm {
		g += math.Pow(ch-0.5, 2.0)
	}
	f, pr := make([]float64, obj), 1.0
	for _, ch := range xc {
		f[0] += pr * math.Cos(math.Pi*math.Pow(ch, alpha)*0.5)
		pr = f[0]
	}
	f[0] *= (1.0 + g)
	j := 1
	for m := obj - 2; m >= 0; m-- {
		pr = 1.0
		for _, ch := range xc[:m] {
			f[j] += pr * math.Cos(math.Pi*math.Pow(ch, alpha)*0.5)
			pr = f[j]
		}
		f[j] *= (1.0 - chrom[m]) * (1.0 + g) * math.Sin(0.5*math.Pow(xc[m], alpha)*math.Pi)
		j++
	}
	return f
}

// DTLZ5 multiobjective function. It returns a slice of obj values.
//
// From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
func DTLZ5(individual *base.Float64Individual, obj int) []float64 {
	chrom := individual.GetChromosome().([]float64)
	gval := 0.0
	for _, a := range chrom[obj-1:] {
		gval += math.Pow(a-0.5, 2.0)
	}

	theta := func(x float64) float64 {
		return math.Pi / (4.0 * (1.0 + gval) * (1.0 + 2.0*gval*x))
	}
	fit, t := make([]float64, obj), 1.0
	for _, a := range chrom[1:] {
		t *= a
	}
	fit[0] = (1 + gval) * math.Cos(math.Pi/2.0*chrom[0]) * t
	j := 1
	for m := obj - 1; m >= 1; m-- {
		if m == 1 {
			fit[j] = (1 + gval) * math.Sin(math.Pi/2.0*chrom[0])
		} else {
			t = 1.0
			for _, a := range chrom[1 : m-1] {
				t *= math.Cos(theta(a))
			}
			t *= math.Sin(theta(chrom[m-1]))
			fit[j] = (1 + gval) * math.Cos(math.Pi/2.0*chrom[0]) * t
		}
	}
	return fit
}

// DTLZ6 multiobjective function. It returns a slice of obj values.
//
// From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
func DTLZ6(individual *base.Float64Individual, obj int) []float64 {
	chrom := individual.GetChromosome().([]float64)
	gval := 0.0
	for _, a := range chrom[obj-1:] {
		gval += math.Pow(a, 0.1)
	}

	theta := func(x float64) float64 {
		return math.Pi / (4.0 * (1.0 + gval) * (1.0 + 2.0*gval*x))
	}
	fit, t := make([]float64, obj), 1.0
	for _, a := range chrom[1:] {
		t *= a
	}
	fit[0] = (1 + gval) * math.Cos(math.Pi/2.0*chrom[0]) * t
	j := 1
	for m := obj - 1; m >= 1; m-- {
		if m == 1 {
			fit[j] = (1 + gval) * math.Sin(math.Pi/2.0*chrom[0])
		} else {
			t = 1.0
			for _, a := range chrom[1 : m-1] {
				t *= math.Cos(theta(a))
			}
			t *= math.Sin(theta(chrom[m-1]))
			fit[j] = (1 + gval) * math.Cos(math.Pi/2.0*chrom[0]) * t
		}
	}
	return fit
}

// DTLZ7 multiobjective function. It returns a slice of obj values.
//
// From: K. Deb, L. Thiele, M. Laumanns and E. Zitzler. Scalable Multi-Objective Optimization Test Problems. CEC 2002, p. 825 - 830, IEEE Press, 2002.
func DTLZ7(individual *base.Float64Individual, obj int) []float64 {
	chrom := individual.GetChromosome().([]float64)
	gval := 0.0
	for _, a := range chrom[obj-1:] {
		gval += a
	}
	gval = 9.0/float64(len(chrom[obj-1:]))*gval + 1.0

	fit, t := make([]float64, obj), 0.0
	for i, a := range chrom[:obj-1] {
		fit[i] = a
	}
	for _, a := range chrom[:obj-1] {
		t += a / (1.0 + gval) * (1.0 + math.Sin(3.0*math.Pi*a))
	}
	fit[obj-1] = (1 + gval) * (float64(obj) - t)
	return fit
}

// Fonseca and Fleming's multiobjective function.
//
// From: C. M. Fonseca and P. J. Fleming, "Multiobjective optimization and multiple constraint handling with evolutionary algorithms -- Part II: Application example", IEEE Transactions on Systems, Man and Cybernetics, 1998.
//
// `f_{\\text{Fonseca}1}(\\mathbf{x}) = 1 - e^{-\\sum_{i=1}^{3}(x_i - \\frac{1}{\\sqrt{3}})^2}`
//
// `f_{\\text{Fonseca}2}(\\mathbf{x}) = 1 - e^{-\\sum_{i=1}^{3}(x_i + \\frac{1}{\\sqrt{3}})^2}`
func Fonseca(individual *base.Float64Individual) []float64 {
	f1, f2 := 0.0, 0.0
	c := 1.0 / math.Sqrt(3)
	for _, ch := range individual.GetChromosome().([]float64)[:3] {
		f1 += math.Pow(ch-c, 2.0)
		f2 += math.Pow(ch+c, 2.0)
	}
	f1 = 1 - math.Exp(-f1)
	f2 = 1 - math.Exp(-f2)
	return []float64{f1, f2}
}

// Poloni is Poloni's multiobjective function on a two attribute *individual*. From: C. Poloni, "Hybrid GA for multi objective aerodynamic shape optimization", in Genetic Algorithms in Engineering and Computer Science, 1997.
//
// `A_1 = 0.5 \\sin (1) - 2 \\cos (1) + \\sin (2) - 1.5 \\cos (2)`
//
// `A_2 = 1.5 \\sin (1) - \\cos (1) + 2 \\sin (2) - 0.5 \\cos (2)`
//
// `B_1 = 0.5 \\sin (x_1) - 2 \\cos (x_1) + \\sin (x_2) - 1.5 \\cos (x_2)`
//
// `B_2 = 1.5 \\sin (x_1) - cos(x_1) + 2 \\sin (x_2) - 0.5 \\cos (x_2)`
//
// `f_{\\text{Poloni}1}(\\mathbf{x}) = 1 + (A_1 - B_1)^2 + (A_2 - B_2)^2`
//
// `f_{\\text{Poloni}2}(\\mathbf{x}) = (x_1 + 3)^2 + (x_2 + 1)^2`
func Poloni(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	x1, x2 := chrom[0], chrom[1]
	a1 := 0.5*math.Sin(1.0) - 2.0*math.Cos(1.0) + math.Sin(2.0) - 1.5*math.Cos(2.0)
	a2 := 1.5*math.Sin(1.0) - math.Cos(1.0) + 2.0*math.Sin(2.0) - 0.5*math.Cos(2.0)
	b1 := 0.5*math.Sin(x1) - 2.0*math.Cos(x1) + math.Sin(x2) - 1.5*math.Cos(x2)
	b2 := 1.5*math.Sin(x1) - math.Cos(x1) + 2.0*math.Sin(x2) - 0.5*math.Cos(x2)
	return []float64{1.0 + math.Pow(a1-b1, 2.0) + math.Pow(a2-b2, 2.0), math.Pow(x1+3.0, 2.0) + math.Pow(x2+1.0, 2.0)}
}

// Dent is test problem Dent which lambda equals to 0.85. Two-objective problem with a "dent". individual has two attributes that take values in [-1.5, 1.5].
// From: Schuetze, O., Laumanns, M., Tantar, E., Coello Coello, C.A., & Talbi, E.-G. (2010).
// Computing gap free Pareto front approximations with stochastic search algorithms.
// Evolutionary Computation, 18(1), 65--96. doi:10.1162/evco.2010.18.1.18103
//
// Note that in that paper Dent source is stated as:
// K. Witting and M. Hessel von Molo. Private communication, 2006.
func Dent(individual *base.Float64Individual) []float64 {
	return DentWithLambda(individual, 0.85)
}

// DentWithLambda is test problem Dent. Two-objective problem with a "dent". individual has two attributes that take values in [-1.5, 1.5].
// From: Schuetze, O., Laumanns, M., Tantar, E., Coello Coello, C.A., & Talbi, E.-G. (2010).
// Computing gap free Pareto front approximations with stochastic search algorithms.
// Evolutionary Computation, 18(1), 65--96. doi:10.1162/evco.2010.18.1.18103
//
// Note that in that paper Dent source is stated as:
// K. Witting and M. Hessel von Molo. Private communication, 2006.
func DentWithLambda(individual *base.Float64Individual, lambda float64) []float64 {
	chrom := individual.GetChromosome().([]float64)
	diff2 := math.Pow(chrom[0]-chrom[1], 2.0)
	sum2 := math.Pow(chrom[0]+chrom[1], 2.0)
	d := lambda * math.Exp(-diff2)
	f1 := 0.5*(math.Sqrt(1.0+diff2)+math.Sqrt(1.0+sum2)+chrom[0]-chrom[1]) + d
	f2 := 0.5*(math.Sqrt(1.0+diff2)+math.Sqrt(1.0+sum2)-chrom[0]+chrom[1]) + d
	return []float64{f1, f2}
}
