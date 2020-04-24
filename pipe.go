package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Pipe is a pipe
type Pipe struct {
	rl.Rectangle
	HasPassedThrough bool
}

const (
	pipeWidth     = 70
	pipeHeight    = height / 2
	pipeBorder    = 3
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
	// TODO: use sprites instead of gradients!
	p := pipe.ToInt32()

	// Top Pipe Border
	bTop := rl.NewRectangle(
		pipe.X-pipeBorder, pipe.Y,
		pipe.Width+pipeBorder*2, pipe.Height+pipeBorder,
	)
	rl.DrawRectangleLinesEx(bTop, pipeBorder, rl.DarkGray)

	// Top Pipe
	gTop := pipe.Rectangle
	rl.DrawRectangleGradientEx(gTop, rl.Lime, rl.Lime, rl.Green, rl.Green)

	// Bottom Pipe Border
	bBottom := rl.NewRectangle(
		pipe.X-pipeBorder, pipe.Y+verticalGap+pipe.Height-pipeBorder,
		pipe.Width+pipeBorder*2, height-pipe.Height+pipeBorder,
	)
	rl.DrawRectangleLinesEx(bBottom, pipeBorder, rl.DarkGray)

	// Bottom Pipe
	gBottom := rl.NewRectangle(
		pipe.X, float32(pipe.Y+verticalGap+pipe.Height),
		pipe.Width, float32(height-p.Height),
	)
	rl.DrawRectangleGradientEx(gBottom, rl.Lime, rl.Lime, rl.Green, rl.Green)
}

func (p *Pipe) CollidesWith(other rl.Rectangle) bool {
	top := rl.NewRectangle(
		p.X-pipeBorder, p.Y,
		p.Width+pipeBorder*2, p.Height+pipeBorder,
	)

	bottom := rl.NewRectangle(
		p.X-pipeBorder, p.Y+verticalGap+p.Height-pipeBorder,
		p.Width+pipeBorder*2, height-p.Height+pipeBorder,
	)

	return rl.CheckCollisionRecs(top, other) || rl.CheckCollisionRecs(bottom, other)
}

func (p *Pipe) IsAround(other rl.Rectangle) bool {
	if p.HasPassedThrough {
		return false
	}

	middle := rl.NewRectangle(
		p.X-pipeBorder, p.Y+p.Height+pipeBorder,
		p.Width,
		verticalGap,
	)

	if rl.CheckCollisionRecs(middle, other) {
		p.HasPassedThrough = true
	}

	return p.HasPassedThrough
}

func (p *Pipe) IsOffscreen() bool {
	return p.X+p.Width+pipeBorder <= 0
}
