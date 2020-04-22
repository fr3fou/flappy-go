package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Pipe is a pipe
type Pipe struct {
	rl.Rectangle
}

const (
	pipeWidth      = 70
	basePipeHeight = Height/2 - 10
)

func NewPipe(x, y int) *Pipe {

	return &Pipe{
		Rectangle: rl.NewRectangle(float32(x), float32(y), pipeWidth, float32(basePipeHeight-rl.GetRandomValue(0, 100))),
	}
}
