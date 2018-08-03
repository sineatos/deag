package pso

import (
	"fmt"

	"github.com/sineatos/deag/base"
)

//Particles is a swarm of particles saving in slice
// type Particles base.Individuals

// Particle is the individual of PSO
type Particle struct {
	base.Float64Individual

	// speed of Particle
	speed []float64
	// pbest
	pbest *Particle
	// minimum of speed
	smin float64
	// maximum of speed
	smax float64
}

// NewParticle returns a Particle
func NewParticle(chromosome, speed []float64, smin, smax float64, fitness *base.Fitness) *Particle {
	return &Particle{
		Float64Individual: *base.NewFloat64Individual(chromosome, fitness),
		speed:             speed,
		smin:              smin,
		smax:              smax,
	}
}

// Clone returns an copy of Particle
func (part *Particle) Clone() interface{} {
	f64Ind := part.Float64Individual.Clone().(*base.Float64Individual)
	cSpeed := make([]float64, len(part.speed))
	copy(cSpeed, part.speed)
	return &Particle{
		Float64Individual: *f64Ind,
		speed:             cSpeed,
		pbest:             part.pbest,
		smin:              part.smin,
		smax:              part.smax,
	}
}

func (part *Particle) String() string {
	fmtStr := "Particle{Float64Individual:%v, speed:%v, smin:%v, smax: %v}"
	return fmt.Sprintf(fmtStr, part.Float64Individual.String(), part.speed, part.smin, part.smax)
}

// GetSpeed returns speed
func (part *Particle) GetSpeed() []float64 {
	return part.speed
}

// SetSpeed sets speed and check if is out of limits
func (part *Particle) SetSpeed(speed []float64) {
	if len(speed) != len(part.speed) {
		panic(fmt.Sprintf("This speed's length is not equals to part.speed's: %v,%v", len(speed), len(part.speed)))
	}
	for i, v := range speed {
		if v < part.smin {
			part.speed[i] = part.smin
		} else if v > part.smax {
			part.speed[i] = part.smax
		} else {
			part.speed[i] = v
		}
	}
}

// GetPBest returns pbest
func (part *Particle) GetPBest() *Particle {
	return part.pbest
}

// SetPBest sets the other as part's pbest
func (part *Particle) SetPBest(other *Particle) {
	part.pbest = other
}
