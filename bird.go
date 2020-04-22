package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Bird is a bird
type Bird struct {
	rl.Rectangle
}

const birdSize = 30

func NewBird(x, y int) *Bird {
	return &Bird{
		Rectangle: rl.NewRectangle(float32(x), float32(y), birdSize, birdSize),
	}
}
