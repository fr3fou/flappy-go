package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Pipe is a pipe
type Pipe struct {
	rl.Rectangle
	Pos rl.Vector2
}

const (
	pipeWidth  = 100
	pipeHeight = Height / 2
)

func NewPipe(x, y float32) *Pipe {

	return &Pipe{
		Pos:       rl.NewVector2(x, y),
		Rectangle: rl.NewRectangle(x, y, pipeWidth, pipeHeight),
	}
}
