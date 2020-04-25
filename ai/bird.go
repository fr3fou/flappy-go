package ai

import (
	"math"

	"github.com/fr3fou/flappy-go/flappy"
	"github.com/fr3fou/gone/gone"
)

type Bird struct {
	flappy.Bird
	Brain *gone.NeuralNetwork
	// Inputs to the brain
	Score   int
	Fitness float64
	Alive   bool
}

func NewBird(x, y int, brain *gone.NeuralNetwork) *Bird {
	return &Bird{
		Bird:  *flappy.NewBird(x, y),
		Alive: true,
		Brain: brain,
	}
}

func (b *Bird) Update() {
	b.Score++
	b.Bird.Update()
}

func (b *Bird) ShouldJump(pipe *flappy.Pipe) bool {
	alpha := float64(b.Y)
	beta := math.Abs(float64(b.X - pipe.X))
	gamma := math.Abs(float64(b.Y - pipe.Height + flappy.PipeBorder))
	delta := math.Abs(float64(b.Y - flappy.Height - pipe.Height + flappy.PipeBorder))

	inputs := []float64{alpha, beta, gamma, delta}
	outputs := b.Brain.Predict(inputs)

	return outputs[0] > outputs[1]
}
