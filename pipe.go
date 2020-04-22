package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Pipe is a pipe
type Pipe struct {
	rl.Rectangle
}

const (
	pipeWidth     = 70
	pipeHeight    = Height / 2
	verticalGap   = 100
	horizontalGap = 150
)

func NewPipe(x, y int) *Pipe {
	return &Pipe{
		Rectangle: rl.NewRectangle(
			float32(x), float32(y),
			pipeWidth, float32(rl.GetRandomValue(verticalGap, pipeHeight)),
		),
	}
}
