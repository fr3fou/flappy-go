package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Pipe is a pipe
type Pipe struct {
	rl.Rectangle
}

const (
	pipeWidth     = 70
	pipeHeight    = Height / 2
	pipeBorder    = 2
	verticalGap   = 120
	horizontalGap = 230
	speed         = 2.3
)

func NewPipe(x, y int) *Pipe {
	return &Pipe{
		Rectangle: rl.NewRectangle(
			float32(x), float32(y),
			pipeWidth, float32(rl.GetRandomValue(verticalGap, pipeHeight)),
		),
	}
}

func (pipe *Pipe) Draw() {
	p := pipe.ToInt32()

	// Top Pipe Border
	rl.DrawRectangle(p.X-pipeBorder, p.Y,
		p.Width+pipeBorder*2, p.Height+pipeBorder,
		rl.DarkGray,
	)
	// Top Pipe
	gTop := pipe.Rectangle
	rl.DrawRectangleGradientEx(gTop, rl.Lime, rl.Lime, rl.Green, rl.Green)

	// Bottom Pipe Border
	rl.DrawRectangle(p.X-pipeBorder, p.Y+verticalGap+p.Height-pipeBorder,
		p.Width+pipeBorder*2, Height-p.Height+pipeBorder,
		rl.DarkGray,
	)
	// Bottom Pipe
	gBottom := rl.NewRectangle(
		pipe.X, float32(pipe.Y+verticalGap+pipe.Height),
		pipe.Width, float32(Height-p.Height),
	)
	rl.DrawRectangleGradientEx(gBottom, rl.Lime, rl.Lime, rl.Green, rl.Green)
}
