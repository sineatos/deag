package benchmarks

import (
	"math"
	"math/rand"

	"github.com/sineatos/deag/base"
)

// Float64Evaluator is the type of float64 evaluator
type Float64Evaluator func(individual *base.Float64Individual) []float64

//*************************************************************
// Unimodal
//*************************************************************

// Rand Random test objective function.
//
// Type: minimization or maximization
//
// Range: none
//
// Global optima: none
//
// Function: `f(x) = f(\mathbf{x}) = \\text{\\texttt{random}}(0,1)``
func Rand(individual *base.Float64Individual) []float64 {
	return []float64{rand.Float64()}
}

// Plane test objective function.
//
// Type: minimization
//
// Range: none
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `f(\mathbf{x}) = x_0`
func Plane(individual *base.Float64Individual) []float64 {
	return []float64{individual.GetChromosome().([]float64)[0]}
}

// Sphere test objective function.
//
// Type: minimization
//
// Range: none
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `x_i = 0, \\forall i \in \\lbrace 1 \\ldots N\\rbrace`, :math:`f(\mathbf{x}) = 0`
func Sphere(individual *base.Float64Individual) []float64 {
	ans := 0.0
	for _, c := range individual.GetChromosome().([]float64) {
		ans += c * c
	}
	return []float64{ans}
}

// Cigar test objective function.
//
// Type: minimization
//
// Range: none
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `x_i = 0, \\forall i \in \\lbrace 1 \\ldots N\\rbrace`, :math:`f(\mathbf{x}) = 0`
func Cigar(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	ans := 0.0
	for _, c := range chrom {
		ans += c
	}
	return []float64{ans*1e6 + chrom[0]}
}

// Rosenbrock test objective function.
//
// Type: minimization
//
// Range: none
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `x_i = 1, \\forall i \in \\lbrace 1 \\ldots N\\rbrace`, :math:`f(\mathbf{x}) = 0`
func Rosenbrock(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	ans := 0.0
	for i, y := range chrom[1:] {
		ans += 100*math.Pow(chrom[i]*chrom[i]-y, 2.0) + math.Pow(1.0-chrom[i], 2.0)
	}
	return []float64{ans}
}

// H1 is simple two-dimensional function containing several local maxima.
//
// From: The Merits of a Parallel Genetic Algorithm in Solving Hard Optimization Problems, A. J. Knoek van Soest and L. J. R. Richard  Casius, J. Biomech. Eng. 125, 141 (2003)
//
// Type: maximization
//
// Range: x_i \in [-100, 100]
//
// Global optima: x=[8.6998, 6.7665],f(x)=1.99999999992158
//
// Function: f(x) = `f(\mathbf{x}) = \\frac{\sin(x_1 - \\frac{x_2}{8})^2 + \\sin(x_2 + \\frac{x_1}{8})^2}{\\sqrt{(x_1 - 8.6998)^2 + (x_2 - 6.7665)^2} + 1}`
func H1(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	num := math.Pow(math.Sin(chrom[0]-chrom[1]/8.0), 2.0) + math.Pow(math.Sin(chrom[1]+chrom[0]/8.0), 2.0)
	deNum := math.Pow(chrom[0]-8.6998, 2.0) + math.Pow(chrom[1]-6.7665, 2)
	deNum = math.Pow(deNum, 0.5) + 1
	return []float64{num / deNum}
}

//*************************************************************
// Multimodal
//*************************************************************

// Ackley test objective function.
//
// Type: minimization
//
// Range: x_i \in [-15, 30]
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `f(\\mathbf{x}) = 20 - 20\exp\left(-0.2\sqrt{\\frac{1}{N} \\sum_{i=1}^N x_i^2} \\right) + e - \\exp\\left(\\frac{1}{N}\sum_{i=1}^N \\cos(2\pi x_i) \\right)`
func Ackley(individual *base.Float64Individual) []float64 {
	n, s, c := individual.Len(), 0.0, 0.0
	chrom := individual.GetChromosome().([]float64)
	for _, ch := range chrom {
		s += ch * ch
		c += math.Cos(2 * math.Pi * ch)
	}
	s = math.Exp(-0.2 * math.Sqrt(1.0/float64(n)*s))
	c = math.Exp(1.0 / float64(n) * c)
	return []float64{20.0 - 20.0*s + math.E - c}
}

// Bohachevsky test objective function.
//
// Ackley test objective function.
//
// Type: minimization
//
// Range: x_i \in [-100, 100]
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `f(\mathbf{x}) = \sum_{i=1}^{N-1}(x_i^2 + 2x_{i+1}^2 - 0.3\cos(3\pi x_i) - 0.4\cos(4\pi x_{i+1}) + 0.7)`
func Bohachevsky(individual *base.Float64Individual) []float64 {
	ans := 0.0
	chrom := individual.GetChromosome().([]float64)
	for i, y := range chrom[1:] {
		ans += chrom[i] * chrom[i]
		ans += 2.0 * y * y
		ans -= 0.3 * math.Cos(3.0*math.Pi*chrom[i])
		ans -= 0.4 * math.Cos(4.0*math.Pi*y)
		ans += 0.7
	}
	return []float64{ans}
}

// Rastrigin test objective function.
//
// Type: minimization
//
// Range: x_i \in [-5.12, 5.12]
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `f(\\mathbf{x}) = 10N + \sum_{i=1}^N x_i^2 - 10 \\cos(2\\pi x_i)`
func Rastrigin(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	ans := 0.0
	for _, c := range chrom {
		ans += c*c - 10*math.Cos(2.0*math.Pi*c)
	}
	return []float64{float64(10*individual.Len()) + ans}
}

// RastriginScaled is Scaled Rastrigin test objective function.
//
// Type: minimization
//
// Function: `f_{\\text{RastScaled}}(\mathbf{x}) = 10N + \sum_{i=1}^N \left(10^{\left(\\frac{i-1}{N-1}\\right)} x_i \\right)^2 x_i)^2 - 10\cos\\left(2\\pi 10^{\left(\\frac{i-1}{N-1}\\right)} x_i \\right)`
func RastriginScaled(individual *base.Float64Individual) []float64 {
	n := individual.Len()
	ans := 10.0 * float64(n)
	for i, c := range individual.GetChromosome().([]float64) {
		t1 := math.Pow(math.Pow(10.0, float64(i/(n-1)))*c, 2.0)
		t2 := 10.0 * math.Cos(2.0*math.Pi*math.Pow(10.0, float64(i/(n-1)))*c)
		ans += t1 - t2
	}
	return []float64{ans}
}

// RastriginSkew is Skewed Rastrigin test objective function.
//
// Type: minimization
//
// Function: `f_{\\text{RastSkew}}(\mathbf{x}) = 10N \sum_{i=1}^N \left(y_i^2 - 10 \\cos(2\\pi x_i)\\right)`
//
// `\\text{with } y_i = \\begin{cases} 10\\cdot x_i & \\text{ if } x_i > 0,\\\ x_i & \\text{ otherwise } \\end{cases}`
func RastriginSkew(individual *base.Float64Individual) []float64 {
	n := individual.Len()
	chrom := individual.GetChromosome().([]float64)
	ans := float64(10 * n)
	for _, c := range chrom {
		if c > 0.0 {
			c *= 10.0
		}
		ans += math.Pow(c, 2.0) - 10.0*math.Cos(2*math.Pi*c)
	}
	return []float64{ans}
}

// Schaffer test objective function.
//
// Type: minimization
//
// Range: x_i \in [-100, 100]
//
// Global optima: x=[0,0,...,0],f(x)=0
//
// Function: `f(\mathbf{x}) = \sum_{i=1}^{N-1} (x_i^2+x_{i+1}^2)^{0.25} \cdot \\left[ \sin^2(50\cdot(x_i^2+x_{i+1}^2)^{0.10}) + 1.0 \\right]`
func Schaffer(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	ans := 0.0
	for i, y := range chrom {
		t1 := math.Pow(chrom[i]*chrom[i]+y*y, 0.25)
		t2 := math.Sin(50.0 * math.Pow(chrom[i]*chrom[i]+y*y, 0.1))
		t2 = math.Pow(t2, 2.0) + 1.0
		ans += t1 * t2
	}
	return []float64{ans}
}

// Schwefel test objective function.
//
// Range: x_i \in [-500, 500]
//
// Global optima: x=[420.96874636,420.96874636,...,420.96874636],f(x)=0
//
// Function: `f(\mathbf{x}) = 418.9828872724339\cdot N - \sum_{i=1}^N\,x_i\sin\\left(\sqrt{|x_i|}\\right)`
func Schwefel(individual *base.Float64Individual) []float64 {
	n := individual.Len()
	ans := 418.9828872724339 * float64(n)
	chrom := individual.GetChromosome().([]float64)
	for _, c := range chrom {
		ans -= c * math.Sin(math.Sqrt(math.Abs(c)))
	}
	return []float64{ans}
}

// Himmelblau function is multimodal with 4 defined minimums in [-6, 6]^2.
//
// Type: minimization
//
// Range: x_i \in [-6, 6]
//
// Global optima: x_1=[3.0,2.0],x_2=[-2.805118, 3.131312],x_3=[-3.779310, -3.283186],x_4=[3.584428, -1.848126],f(x)=0
//
// Function: `f(x_1, x_2) = (x_1^2 + x_2 - 11)^2 + (x_1 + x_2^2 -7)^2`
func Himmelblau(individual *base.Float64Individual) []float64 {
	chrom := individual.GetChromosome().([]float64)
	return []float64{
		math.Pow(chrom[0]*chrom[0]+chrom[1]-11.0, 2.0) + math.Pow(chrom[0]+chrom[1]*chrom[1]-7.0, 2.0),
	}
}

// Shekel multimodal function can have any number of maxima. The number of maxima is given by the length of any of the arguments a or c, a is a matrix of size :math:M * N, where M is the number of maxima and *N* the number of dimensions and c is a M * 1 vector.
//
// `f_\\text{Shekel}(\mathbf{x}) = \\sum_{i = 1}^{M} \\frac{1}{c_{i} + \\sum_{j = 1}^{N} (x_{j} - a_{ij})^2 }`
//
// The following figure uses
//
// `\\mathcal{A} = \\begin{bmatrix} 0.5 & 0.5 \\\\ 0.25 & 0.25 \\\\  0.25 & 0.75 \\\\ 0.75 & 0.25 \\\\ 0.75 & 0.75 \\end{bmatrix}` and :math:`\\mathbf{c} = \\begin{bmatrix} 0.002 \\\\ 0.005 \\\\ 0.005 \\\\ 0.005 \\\\ 0.005 \\end{bmatrix}`, thus defining 5 maximums in :math:`\\mathbb{R}^2`
func Shekel(individual *base.Float64Individual, a [][]float64, c []float64) []float64 {
	chrom := individual.GetChromosome().([]float64)
	ans := 0.0
	for i := range c {
		t := 0.0
		for j, aij := range a[i] {
			t += math.Pow(chrom[j]-aij, 2)
		}
		ans += 1.0 / t
	}
	return []float64{ans}
}
