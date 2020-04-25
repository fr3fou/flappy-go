package ai

import (
	"math"

	"github.com/fr3fou/flappy-go/flappy"
	"github.com/fr3fou/gone/gone"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bird struct {
	flappy.Bird
	Brain   *gone.NeuralNetwork
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

func (b *Bird) Draw(opacity uint8) {
	if b.Alive {
		b.Bird.Draw(rl.NewColor(rl.Yellow.R, rl.Yellow.G, rl.Yellow.B, opacity))
	} else {
		b.Bird.Draw(rl.NewColor(rl.Red.R, rl.Red.G, rl.Red.B, opacity))
	}
}
