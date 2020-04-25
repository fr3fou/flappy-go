package ai

import (
	"github.com/fr3fou/flappy-go/flappy"
	"github.com/fr3fou/gone/gone"
)

type Bird struct {
	flappy.Bird
	Brain *gone.NeuralNetwork
}

func NewBird(x, y int, brain *gone.NeuralNetwork) *Bird {
	return &Bird{
		Bird:  *flappy.NewBird(x, y),
		Brain: brain,
	}
}
